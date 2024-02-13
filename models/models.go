package models

import "gorm.io/gorm"

//! User models will store user data
type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Email    string `gorm:"unique;email"`
	Password string `gorm:"password"`
}

//! Questions struct will store the question and its options
type Questions struct {
	Category      string `json:"category"`
	Question      string `json:"question"`
	Option1       string `json:"option1"`
	Option2       string `json:"option2"`
	Option3       string `json:"option3"`
	Option4       string `json:"option4"`
	CorrectOption string `json:"correct_option"`
}

//! Golang struct is table which will store all golang questions
type Golang struct {
	gorm.Model
	Question      string `json:"question"`
	Option1       string `json:"option1"`
	Option2       string `json:"option2"`
	Option3       string `json:"option3"`
	Option4       string `json:"option4"`
	CorrectOption string `json:"correct_option"`
}

//! Python struct is table which will store all python questions
type Python struct {
	gorm.Model
	Question      string `json:"question"`
	Option1       string `json:"option1"`
	Option2       string `json:"option2"`
	Option3       string `json:"option3"`
	Option4       string `json:"option4"`
	CorrectOption string `json:"correct_option"`
}

//! Javascript struct is table which will store all javascript questions
type Javascript struct {
	gorm.Model
	Question      string `json:"question"`
	Option1       string `json:"option1"`
	Option2       string `json:"option2"`
	Option3       string `json:"option3"`
	Option4       string `json:"option4"`
	CorrectOption string `json:"correct_option"`
}
