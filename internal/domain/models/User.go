package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"-" gorm:"PrimaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	Email     string    `json:"email,omitempty" gorm:"unique;not null" validate:"required,email"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
}
