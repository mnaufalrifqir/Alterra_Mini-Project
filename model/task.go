package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	UserID       uint
	Name         string    `json:"name" form:"name" validate:"required"`
	Status       string    `json:"status" gorm:"type:enum('finished', 'unfinished');default:'unfinished';not_null"`
	DueDate      time.Time `json:"due_date" form:"due_date" validate:"required"`
	TaskPriority string    `json:"task_priority" gorm:"type:enum('high', 'medium', 'low');default:'low';not_null"`
	Description  string    `json:"description" form:"description"`
	User         User      `json:"-" gorm:"foreignKey:UserID" validate:"-"`
}
