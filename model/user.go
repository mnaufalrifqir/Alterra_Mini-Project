package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password"`
	Task     []Task `gorm:"foreignkey:UserID"`
}
