package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" gorm:"unique" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Task     []Task `gorm:"foreignkey:UserID"`
}
