package helpers

import (
	"log"

	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
)

// ! Golang function will add golang questions
func Golang(data models.Questions) error {
	questions := models.Golang{}
	questions.Question = data.Question
	questions.Option1 = data.Option1
	questions.Option2 = data.Option2
	questions.Option3 = data.Option3
	questions.Option4 = data.Option4
	questions.CorrectOption = data.CorrectOption

	DB, err := database.ConnectToDatabase()

	if err != nil {
		log.Println("Error connecting to database in Golang function : ", err)
		return err
	}

	result := DB.Create(&questions)

	if result.Error != nil {
		log.Println("Error in saving the question of particular category ", result.Error)
		return result.Error
	}
	return nil

}

//! Python function will add python questions
func Python(data models.Questions) error {
	questions := models.Python{}
	questions.Question = data.Question
	questions.Option1 = data.Option1
	questions.Option2 = data.Option2
	questions.Option3 = data.Option3
	questions.Option4 = data.Option4
	questions.CorrectOption = data.CorrectOption

	DB, err := database.ConnectToDatabase()

	if err != nil {
		log.Println("Error connecting to database in Python function : ", err)
		return err
	}

	result := DB.Create(&questions)

	if result.Error != nil {
		log.Println("Error in saving the question of particular category ", result.Error)
		return result.Error
	}
	return nil

}
