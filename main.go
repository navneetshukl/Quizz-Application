package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/auth"
	"github.com/navneetshukl/database"
)

func init() {
	database.MigrateDatabase()
}
func main() {
	router := gin.Default()

	router.POST("/api/register", auth.Register)

	router.Run()

}
