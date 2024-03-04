package main

import (
	"project/controllers"
	"project/database"
	"project/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)

	authorized := r.Group("/")
	authorized.Use(middlewares.Authenticate())
	{
		authorized.POST("/photos", controllers.UploadPhoto)
		authorized.DELETE("/photos/:photoId", controllers.DeletePhoto)
	}

	r.Run(":8080")
}
