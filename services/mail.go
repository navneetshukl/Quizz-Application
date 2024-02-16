package services

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
	"github.com/navneetshukl/models"
)

// ! SendMail function will send the score of user to his mail
func SendMail(email, name string, data models.Mail) error {
	err := godotenv.Load()

	if err != nil {
		log.Println("Error loading .env file ", err)
		return err
	}

	mail_password := os.Getenv("MAIL_PASSWORD")

	auth := smtp.PlainAuth("", "yatinjal123@gmail.com", mail_password, "smtp.gmail.com")

	msg := fmt.Sprintf("Hello %s . Your score in the recent test in %s is %d out of %d .", name, data.Subject, data.Total, data.Maximum)
	emails := []string{email}

	err = smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"yatinjal123@gmail.com",
		emails,
		[]byte(msg),
	)

	if err != nil {
		log.Println("Error in sending the mail ", err)
		return err
	}

	return nil
}
