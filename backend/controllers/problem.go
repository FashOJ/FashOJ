package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
	"github.com/gin-gonic/gin"
)

func CreateOrUpdataProblem(ctx *gin.Context) {
	if ctx.Value("user").(models.User).Permission != global.AdminUser {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "you don't have right to create or updata problem"})
		return
	}

	var problem models.Problem
	if err := ctx.ShouldBindJSON(&problem); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	problem.Author = ctx.Value("user").(models.User)

	if err := global.DB.AutoMigrate(&models.Problem{}, &models.Example{}, &models.Limit{}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := global.DB.Save(&problem).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func UploadTestcase(ctx *gin.Context) {
	if err := global.DB.AutoMigrate(&models.Testcase{}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tempDir := os.TempDir()
	filePath := path.Join(tempDir, file.Filename)
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	zipFile, err := zip.OpenReader(filePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer zipFile.Close()

	problemID := ctx.Param("pid")
	if problemID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "problem ID is required"})
		return
	}

	inputs := make(map[string]*zip.File)
	outputs := make(map[string]*zip.File)

	var testcases []models.Testcase
	for _, file := range zipFile.File {
		if !isTestcaseFile(file.Name) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "文件格式错误"})
			return
		}
		isInput := strings.HasSuffix(file.Name, ".in")
		if isInput {
			inputs[strings.TrimSuffix(file.Name, ".in")] = file
		} else {
			outputs[strings.TrimSuffix(file.Name, ".out")] = file
		}
	}

	for name := range inputs {
		infile, err := inputs[name].Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer infile.Close()

		output, ok := outputs[name]
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "缺少对应的输出文件: " + name + ".out"})
			return
		}

		outfile, err := output.Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		defer outfile.Close()

		incontent, err := io.ReadAll(infile)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		outcontent, err := io.ReadAll(outfile)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		inputFileName := fmt.Sprintf("%s_%s.in", problemID, name)
		outputFileName := fmt.Sprintf("%s_%s.out", problemID, name)

		inputFilePath := path.Join("file", inputFileName)
		outputFilePath := path.Join("file", outputFileName)

		if err := os.MkdirAll("file", os.ModePerm); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := saveFile(bytes.NewReader(incontent), inputFilePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		if err := saveFile(bytes.NewReader(outcontent), outputFilePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		testcases = append(testcases, models.Testcase{
			ProblemID: problemID,
			Input:     inputFileName,
			Output:    outputFileName,
		})
	}

	var problem models.Problem
	if err := global.DB.Where("problem_id = ?", problemID).First(&problem).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	problem.Testcase = testcases
	if err := global.DB.Save(&problem).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := os.Remove(filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}

func saveFile(reader io.Reader, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	return err
}

func isTestcaseFile(filename string) bool {
	matched, _ := regexp.MatchString(`^\d+\.(in|out)$`, filename)
	return matched
}
