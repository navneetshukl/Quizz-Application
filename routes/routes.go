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

//! GetQuestionsRoute function will get the questions from database
func GetQuestionsRoute(c *gin.Context) {
	cat := c.Param("cat")

	if cat == "golang" {
		err, data := helpers.GolangQuestions()

		if err != nil {
			log.Println("Error in getting the golang question : ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error in getting the golang question",
			})
		}
		c.JSON(http.StatusOK, data)
	} else if cat == "python" {

		err, data := helpers.PythonQuestions()
		if err != nil {
			log.Println("Error in getting the python question : ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error in getting the python question",
			})
		}
		c.JSON(http.StatusOK, data)

	} else if cat == "javascript" {
		err, data := helpers.JavascriptQuestions()
		if err != nil {
			log.Println("Error in getting the javascript question : ", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error in getting the javascript question",
			})
		}
		c.JSON(http.StatusOK, data)

	}
}
