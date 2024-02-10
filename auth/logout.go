package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	log.Println("I am on signup page")

	c.JSON(http.StatusOK, gin.H{
		"message": "User Logout successfully",
	})

}
