package helpers

import (
	"fmt"
	"log"

	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
)

func StoreQuestion(data models.Questions) error {

	if data.Category == "Python" {
		err := Python(data)
		if err != nil {
			log.Println("Error in storing the python question : ", err)
			return err
		}
	} else if data.Category == "Golang" {
		err := Golang(data)
		if err != nil {
			log.Println("Error in storing the golang question : ", err)
			return err
		}
	} else if data.Category == "Javascript" {

		err := Javascript(data)
		if err != nil {
			log.Println("Error in storing the javascript question : ", err)
			return err
		}
	}

	return nil

}

// ! Getname function will return name of signed in user
func Getname(email string) (string, error) {

	var data models.User

	DB, err := database.ConnectToDatabase()
	if err != nil {
		log.Println("Error in connecting to database ", err)
		return "", err
	}

	result := DB.Where("email=?", email).First(&data)
	if result.Error != nil {
		log.Println("Error in getting the data from database ", err)
		return "", result.Error
	}
	fmt.Println("Data issss ", data)
	return data.Name, nil
}
