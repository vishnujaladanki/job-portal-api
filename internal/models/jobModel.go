package models

import "gorm.io/gorm"

type Job struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	CompanyID   uint // Foreign key for the Company
}

type NewJob struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	CompanyID   uint
}
