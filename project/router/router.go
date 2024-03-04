package main

import (
	"project/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/users/register", controllers.RegisterUser)
	r.GET("/users/login", controllers.LoginUser)

	r.Run() // listen and serve on 0.0.0.0:8080
}
