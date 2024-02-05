package models

import "gorm.io/gorm"

//! User models will store user data
type User struct {
	gorm.Model
	Name     string `gorm:"name"`
	Email    string `gorm:"unique;email"`
	Password string `gorm:"password"`
}


//! Admin models will store admin data
type Admin struct {
	gorm.Model
	Name     string `gorm:"name"`
	Email    string `gorm:"unique;email"`
	Password string `gorm:"password"`
}
