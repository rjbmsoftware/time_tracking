package main

import (
	"log"
	docs "time-tracking/docs"
	"time-tracking/internal/controllers"
	"time-tracking/internal/models"
	"time-tracking/internal/routers"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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

	db.AutoMigrate(&models.Project{}, &models.User{})

	if !controllers.AdminUserExists(db) {
		log.Println("No admin user found")
		password, adminUser := controllers.CreateDefaultAdminUser(db)
		message := "Admin user created, username: %s, password: %s, email: %s"
		log.Printf(message, adminUser.Name, password, adminUser.Email)
	}

	// db.Create(&Project{Name: "test"})
	// db.Create(&Project{Name: "another test"})
	// router := gin.Default()
	validate := validator.New()
	authController := controllers.NewAuthControllerImpl(db, validate)
	router := routers.AuthRouter(authController)
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/projects", controllers.GetProjects)
	router.GET("/projects/:id", controllers.GetProject)
	router.POST("/projects", controllers.PostProjects)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("localhost:8080")
}
