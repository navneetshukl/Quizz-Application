package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// ! Authenticate function is middleware function which will authenticate the JWT token
func Authenticate(c *gin.Context) {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error in loading the .env file ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in loading the .env",
		})

		return
	}

	tokenString, err := c.Cookie("Authorization")
	secret := os.Getenv("SECRET")

	if err != nil {
		log.Println("Error in Getting the tokenstring from Cookie ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in Getting the tokenstring from Cookie",
		})

	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {

			log.Println("Token is expired")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Token is expired",
			})
			return
		}

		email := claims["sub"].(string)

		if email == "" {
			log.Println("No email in the token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "No email in the token",
			})
			return
		}
		c.Set("user", email)
		c.Next()
	} else {

		log.Println("Error in parsing token")
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Error in parsing token",
		})
	}

}
