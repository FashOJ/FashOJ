package controllers

import (
	"FashOJ/Judger/internal/dto"
	"FashOJ/Judger/internal/global"
	"FashOJ/Judger/internal/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProblem(c *gin.Context) {

	var problem models.Problem
	var req dto.CreateProblem

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err := global.DB.Where("problem_id = ?", req.ProblemId).First(&problem).Error; err != nil {

		if !errors.Is(err, gorm.ErrRecordNotFound) {

			c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
			return
		}
	}

	problem.ProblemId = req.ProblemId
	problem.TimeLimit = req.TimeLimit
	problem.MemoryLimit = req.MemoryLimit

	if err := global.DB.Save(&problem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "internal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}