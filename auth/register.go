package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
	"golang.org/x/crypto/bcrypt"
)

//!  Register function will register the user
func Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error in encrypting the password in Register function ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in encrypting the password",
		})
		return
	}
	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(encryptedPassword),
	}

	DB, err := database.ConnectToDatabase()

	if err != nil {
		log.Println("Error in connecting to Database in Register function", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in connecting to Database",
		})
		return

	}

	result := DB.Create(&user)
	if result.Error != nil {
		log.Println("Error in saving to Database in Register function ", result.Error)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Eror in saving to Database ",
			"error":   result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Registered Successfully",
	})

}

//! RegisterForm function will render the register form
func RegisterForm(c *gin.Context) {
	c.HTML(http.StatusOK, "register.tmpl", nil)
}
