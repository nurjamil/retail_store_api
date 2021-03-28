package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" form:"username"`
	Name string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	PhoneNumber string `json:"phoneNumber" form:"phoneNumber"`
	Token string `json:"token" form:"token"`
}
