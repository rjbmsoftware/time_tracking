package main

import (
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

	// TODO: setup initial admin user if one does not exist

	db.AutoMigrate(&Project{}, &User{})
	adminUserExists(db)
	// db.Create(&Project{Name: "test"})
	// db.Create(&Project{Name: "another test"})

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
