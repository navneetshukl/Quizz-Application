package routes

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/helpers"
	"github.com/navneetshukl/models"
)

// ! AddQuestionRoute function will add the question
func AddQuestionRoute(c *gin.Context) {
	requestBody := models.Questions{}

	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		log.Println("Error in reading request body in AddQuestionRoute function:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in reading request body",
		})
		return
	}

	err = helpers.StoreQuestion(requestBody)

	if err != nil {
		log.Println("Error in storing the question : ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in storing the question",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "question stored successfully",
	})

}
