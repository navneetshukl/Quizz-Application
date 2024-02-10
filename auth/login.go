package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {

	var requestData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	//? Read the request body

	err := c.ShouldBindJSON(&requestData)

	if err != nil {
		log.Println("Error in reading request body in login function:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in reading request body",
		})
		return
	}

	db, err := database.ConnectToDatabase()

	if err != nil {
		log.Println("Error in connecting to database in login function:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in connecting to database",
		})
		return
	}

	var loginData models.User

	db.Where("email=?", requestData.Email).First(&loginData)
	if loginData.ID == 0 {
		log.Println("Email Does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email Does not Exist",
		})
		return
	}

	//? Compare the password using bcrypt

	err = bcrypt.CompareHashAndPassword([]byte(loginData.Password), []byte(requestData.Password))
	if err != nil {
		if err != nil {
			log.Println("Password does not exist")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Password does not match",
			})
			return
		}
	}
	//* Implement the JWT here

	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": loginData.Email,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println("Error in saving the JWT Token ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error in saving the JWT Token",
		})
		return
	}

	//* Saving the JWT token to the cookies

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, int(time.Hour*24*30), "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "User Login successfully",
	})

}
