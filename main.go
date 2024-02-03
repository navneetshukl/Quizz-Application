package main

import (
	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/routes"
)

func main() {
	router := gin.Default()

	router.GET("/", routes.Home)
	router.Run()

}
