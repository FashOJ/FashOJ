package controllers

import (
	"FashOJ_Backend/global"
	"FashOJ_Backend/models"
	"FashOJ_Backend/permission"
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

// CreateProblem creates a new problem in the database.
// It requires the user to have permission to create or update problems.
// The problem data is expected to be provided in the request body as JSON.
// If the problem is successfully created, it returns a 200 OK response.
// If there is an error, it returns an appropriate error response.
func CreateProblem(ctx *gin.Context) {

	// Check if the user has permission to create or update problems
	if permission.HasPermission(ctx.Value("user").(models.User), permission.CreateProblem) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You don't have permission to create",
		})
		return
	}

	// Bind the problem data from the request body.
	var problem models.Problem
	if err := ctx.ShouldBindJSON(&problem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Set the author of the problem to the current user.
	problem.Author = ctx.Value("user").(models.User)

	// Save the problem to the database.
	if err := global.DB.Save(&problem).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// ModifyProblem modifies an existing problem in the database.
// It requires the user to have permission to create or update problems.
// The problem ID should be provided in the URL parameter, and the updated problem data is expected to be provided in the request body as JSON.
// If the problem is successfully modified, it returns a 200 OK response.
// If there is an error, it returns an appropriate error response.
func ModifyProblem(ctx *gin.Context) {
	// Check if the user has permission to create or update problems
	// if ctx.Value("user").(models.User).Permission != global.AdminUser {
	if permission.HasPermission(ctx.Value("user").(models.User), permission.ModifyProblem) {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"message": "You don't have permission to modify problem",
		})
		return
	}

	// Get the problem ID from the URL parameter
	problemID := ctx.Param("pid")
	if problemID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "You don't have a valid Problem ID",
		})
		return
	}

	// Find the existing problem in the database
	var problem models.Problem
	if err := global.DB.Where("problem_id = ?", problemID).First(&problem).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Problem not found",
		})
		return
	}

	// Bind the updated problem data from the request body
	var updatedProblem models.Problem
	if err := ctx.ShouldBindJSON(&updatedProblem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Update the problem fields
	problem.Title = updatedProblem.Title
	problem.Content = updatedProblem.Content
	problem.Difficulty = updatedProblem.Difficulty
	problem.Author = ctx.Value("user").(models.User)
	problem.AuthorID = problem.Author.ID
	problem.Limit = updatedProblem.Limit

	// Save the updated problem to the database
	if err := global.DB.Save(&problem).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func UploadTestcase(ctx *gin.Context) {

	// Get the file from the request.
	uploadedFile, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Save the file(Uploaded ZIP TestCase) to the temporary folder.
	tempZipFilePath := path.Join(global.SystemTempFolder, uploadedFile.Filename)
	if err := ctx.SaveUploadedFile(uploadedFile, tempZipFilePath); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Open the ZIP file of the test cases.
	zipFile, err := zip.OpenReader(tempZipFilePath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	defer zipFile.Close()

	// Get the problem ID from the URL parameter.
	problemID := ctx.Param("pid")
	if problemID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Problem ID is required",
		})
		return
	}

	// Classify the test cases into input and output files.
	// Input files should have a ".in" extension, and output files should have a ".out" extension.
	testCasesInputs := make(map[string]*zip.File)
	testCaseOutputs := make(map[string]*zip.File)
	for _, file := range zipFile.File {
		if !isTestcaseFile(file.Name) {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "File format error: " + file.Name + "\nShould be like: 1.in 1.out 2.in 2.out ...",
			})
			return
		}
		isInput := strings.HasSuffix(file.Name, ".in")
		if isInput {
			testCasesInputs[strings.TrimSuffix(file.Name, ".in")] = file
		} else {
			testCaseOutputs[strings.TrimSuffix(file.Name, ".out")] = file
		}
	}

	var testcases []models.Testcase

	// Process each input file and its corresponding output file.
	// For each input file, read its content and save it to a file with the format "problemID_inputFileName".
	// For each output file, read its content and save it to a file with the format "problemID_outputFileName".
	// Create a Testcase object with the problem ID, input file name, and output file name.
	// Append the Testcase object to the testcases slice.
	for name := range testCasesInputs {
		infile, err := testCasesInputs[name].Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer infile.Close()

		// Check if the corresponding output file exists.
		output, ok := testCaseOutputs[name]
		if !ok {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "The corresponding output file is missing: " + name + ".out",
			})
			return
		}
		outfile, err := output.Open()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		defer outfile.Close()

		// Set the file name and save the input and output files.
		// If the file directory does not exist, create it.
		inputFileName := fmt.Sprintf("%s_%s.in", problemID, name)
		outputFileName := fmt.Sprintf("%s_%s.out", problemID, name)
		inputFilePath := path.Join("file", inputFileName)
		outputFilePath := path.Join("file", outputFileName)
		if err := os.MkdirAll("file", os.ModePerm); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Read the input and output file content.
		inTestCaseContent, err := io.ReadAll(infile)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		outTestCaseContent, err := io.ReadAll(outfile)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Save the input and output files to the file directory.
		if err := saveFile(bytes.NewReader(inTestCaseContent), inputFilePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
		if err := saveFile(bytes.NewReader(outTestCaseContent), outputFilePath); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		// Create a Testcase object and append it to the testcases slice.
		// The Testcase object contains the problem ID, input file name, and output file name.
		testcases = append(testcases, models.Testcase{
			ProblemID: problemID,
			Input:     inputFileName,
			Output:    outputFileName,
		})
	}

	// Find the problem in the database and update its test cases.
	// Upgrade the problem object and save it to the database.
	var problem models.Problem
	if err := global.DB.Where("problem_id = ?", problemID).First(&problem).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	problem.Testcase = testcases
	if err := global.DB.Save(&problem).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Delete the temporary ZIP file.
	/*
		If the temporary ZIP file is not deleted, it will occupy disk space and can cause problems when uploading new test cases.
		Therefore, it is important to delete the temporary ZIP file after the test cases have been processed.
	*/
	if err := os.Remove(tempZipFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

// DownloadTestcase downloads the test cases of a problem in ZIP format.
// It requires the user to have permission to view problems.
func saveFile(reader io.Reader, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, reader)
	return err
}

// isTestcaseFile checks if the given filename is a valid test case file.
// A valid test case file should have a ".in" or ".out" extension.
func isTestcaseFile(filename string) bool {
	matchedFileName, _ := regexp.MatchString(`^\d+\.(in|out)$`, filename)
	return matchedFileName
}
