package main

import (
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

	router.POST("/api/register", auth.Register)
	router.POST("/api/login", auth.Login)
	router.GET("/api/logout", middleware.Authenticate, auth.Logout)

	router.POST("/add/question", middleware.Authenticate, routes.AddQuestionRoute)

	router.Run()

}
