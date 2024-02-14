package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/auth"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/middleware"
	"github.com/navneetshukl/routes"
)

func init() {
	database.MigrateDatabase()

}

func main() {
	router := gin.Default()
	
	//? Apply CORS middleware

	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOrigins = []string{"http://localhost:5173"}
	router.Use(cors.New(config))

	router.POST("/api/register", auth.Register)
	router.POST("/api/login", auth.Login)
	router.GET("/api/logout", middleware.Authenticate, auth.Logout)

	router.POST("/add/question", middleware.Authenticate, routes.AddQuestionRoute)

	router.GET("/test/:cat", middleware.Authenticate, routes.GetQuestionsRoute)

	router.Run()

}
