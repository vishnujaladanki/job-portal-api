package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId       uint   `gorm:"primaryKey;autoIncrement"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"-"`
}

type NewUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type Reset struct {
	Otp             string `json:"otp" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	NewPassword     string `json:"new_password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}
