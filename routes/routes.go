package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/helpers"
	"github.com/navneetshukl/models"
	"github.com/navneetshukl/services"
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

// ! GetQuestionsRoute function will get the questions from database
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

// ! SendMailRoute function will send the mail
func SendMailRoute(c *gin.Context) {

	email, ok := c.Get("user")

	if !ok {
		log.Println("Unable to get the email ")
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please try after again",
		})

		return
	}

	var data models.Mail

	err := c.ShouldBindJSON(&data)

	if err != nil {
		log.Println("Error in reading the body of the request for score of test ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error in reading the body",
		})
		return
	}

	name, err := helpers.Getname(email.(string))
	if err != nil {
		log.Println("Error in getting the name ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please try again",
		})
		return
	}

	fmt.Println("Name is ", name)

	helpers.SaveScore(email.(string), data)

	err = services.SendMail(email.(string), name, data)
	if err != nil {
		log.Println("Error in sending the body ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some error occured.Please retry again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": "Mail sent successfully",
	})

}

// ! GetDetailsRoute function will get the details of user for every test
func GetDetailsRoute(c *gin.Context) {

	email, ok := c.Get("user")
	if !ok {
		log.Println("Unable to get the user's email ")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some error occured.Please retry",
		})
		return
	}

	
	DB, err := database.ConnectToDatabase()
	if err != nil {
		log.Println("Error in connecting to database ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some error occured.Please retry",
		})

		return
	}

	data := []models.Score{}
	//DB.Select("total","subject","maximum")

	result := DB.Where("email=?", email.(string)).Select("total", "subject", "maximum").Find(&data)

	if result.Error != nil {
		log.Println("Error in getting the users data from database")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some Error Occured.Please retry",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}
