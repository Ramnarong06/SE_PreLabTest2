package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username		string 		`gorm:"uniqueIndex" valid:"required~Username is required"`
	Password		string
	Email			string 		`gorm:"uniqueIndex" valid:"email~Email is invalid"`
}