package main

import (
	"jwt-go/controllers"
	"jwt-go/initializers"
	"jwt-go/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
