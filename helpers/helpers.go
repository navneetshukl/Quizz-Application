package helpers

import (
	"log"

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
	}

	return nil

}
