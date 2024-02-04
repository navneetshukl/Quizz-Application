package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/routes"
)

func init() {
	database.MigrateDatabase()
}
func main() {
	database.MigrateDatabase()
	router := gin.Default()

	router.GET("/", routes.Home)
	router.Run()

}
