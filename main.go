package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/projects", getProjects)
	router.GET("/projects/:id", getProject)
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

func getProject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be an integer"})
		return
	}

	for _, proj := range projects {
		if proj.ID == id {
			c.IndentedJSON(http.StatusOK, proj)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "project not found"})
}

func postProjects(c *gin.Context) {
	var newProject project

	if err := c.BindJSON(&newProject); err != nil {
		return
	}

	projects = append(projects, newProject)
	c.IndentedJSON(http.StatusCreated, newProject)
}
