package helpers

import (
	"log"

	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
)

//! SaveScore function will save score to database
func SaveScore(email string, data models.Mail) error {

	DB, err := database.ConnectToDatabase()
	if err != nil {
		log.Println("Error in connecting to database ", err)
		return err
	}

	var score models.Score
	score.Email = email
	score.Maximum = data.Maximum
	score.Subject = data.Subject
	score.Total = data.Total

	result := DB.Create(&score)

	if result.Error != nil {
		log.Println("Error in saving to database ", result.Error)
		return result.Error
	}
	return nil

}
