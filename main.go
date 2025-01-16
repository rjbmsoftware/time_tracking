package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/projects", getProjects)
	router.POST("/projects", postProjects)

	router.Run("localhost:8080")
}

type project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var projects = []project{
	{ID: 1, Name: "mobile"},
	{ID: 2, Name: "web"},
	{ID: 3, Name: "admin"},
}

func getProjects(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, projects)
}

func postProjects(c *gin.Context) {
	var newProject project

	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}
