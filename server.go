package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jetsadawwts/go-jwt/controllers"
	"github.com/jetsadawwts/go-jwt/initializers"
	"github.com/jetsadawwts/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run()	
}