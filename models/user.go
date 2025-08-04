package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"  binding:"required" gorm:"type:varchar(255)"`
	Email string `json:"email" binding:"required,email" gorm:"type:varchar(255);unique"`
}
