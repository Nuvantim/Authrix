package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"PrimaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"not null" validate:"required"`
	Email     string    `json:"email,omitempty" gorm:"unique;not null" validate:"required,email"`
	Password  string    `json:"-" gorm:"not null"`
	Roles     []Role   `json:"roles" gorm:"many2many:user_roles;"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
}
