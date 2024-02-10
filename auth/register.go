package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
	"golang.org/x/crypto/bcrypt"
)

// !  Register function will register the user
func Register(c *gin.Context) {
	var requestBody struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//? Bind the request body to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println("Error in reading request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request format",
		})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)

	if err != nil {
		log.Println("Error in encrypting the password in Register function ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in encrypting the password",
		})
		return
	}
	user := models.User{
		Name:     requestBody.Email,
		Email:    requestBody.Email,
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
