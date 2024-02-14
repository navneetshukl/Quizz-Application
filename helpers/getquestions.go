package helpers

import (
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
)

// !  GolangQuestions function will get the golang questions from database
func GolangQuestions() (error, []models.Golang) {

	var data []models.Golang

	DB, err := database.ConnectToDatabase()
	if err != nil {
		return err, data
	}

	result := DB.Select("question", "option1", "option2", "option3", "option4", "correct_option").Limit(10).Find(&data)

	if result.Error != nil {
		return result.Error, data
	}
	return nil, data

}

// !  PythonQuestions function will get the python questions from database
func PythonQuestions() (error, []models.Python) {

	var data []models.Python

	DB, err := database.ConnectToDatabase()
	if err != nil {
		return err, data
	}

	result := DB.Select("question", "option1", "option2", "option3", "option4", "correct_option").Limit(10).Find(&data)

	if result.Error != nil {
		return result.Error, data
	}
	return nil, data

}

// !  JavascriptQuestions function will get the javascript questions from database
func JavascriptQuestions() (error, []models.Javascript) {

	var data []models.Javascript

	DB, err := database.ConnectToDatabase()
	if err != nil {
		return err, data
	}

	result := DB.Select("question", "option1", "option2", "option3", "option4", "correct_option").Limit(10).Find(&data)

	if result.Error != nil {
		return result.Error, data
	}
	return nil, data

}
