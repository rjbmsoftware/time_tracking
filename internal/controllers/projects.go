package controllers

import (
	"net/http"
	"strconv"
	"time-tracking/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func getProjects(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var projects []models.Project
	result := db.Find(&projects)

	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "something went wrong"})
	}
	c.IndentedJSON(http.StatusOK, projects)
}

func getProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be an integer"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)

	var project models.Project
	if err := db.First(&project, "id = ?", id).Error; err == nil {
		c.IndentedJSON(http.StatusOK, project)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
}

func postProjects(c *gin.Context) {
	var newProject models.Project

	if err := c.BindJSON(&newProject); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error"})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	result := db.Create(&newProject)
	if result.Error != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error})
		return
	}

	c.IndentedJSON(http.StatusCreated, newProject)
}
