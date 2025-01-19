package main

import (
	"net/http"
	"strconv"

	docs "time-tracking/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	docs.SwaggerInfo.BasePath = "/api/v1"
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Project{})
	db.Create(&Project{Name: "test"})
	db.Create(&Project{Name: "another test"})

	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/projects", getProjects)
	router.GET("/projects/:id", getProject)
	router.POST("/projects", postProjects)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}

type Project struct {
	gorm.Model
	Name string `json:"name"`
}

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
	var projects []Project
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

	var project Project
	if err := db.First(&project, "id = ?", id).Error; err == nil {
		c.IndentedJSON(http.StatusOK, project)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
}

func postProjects(c *gin.Context) {
	var newProject Project

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
