package models

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"unique"`
	Location string

	//Users []User // Relationship: A company can have multiple users
	Jobs []Job `json:"-"` // Relationship: A company can have multiple jobs
}

type NewCompany struct {
	Name     string `gorm:"unique" ;json:"name" ;validate:"required"`
	Location string `json:"location" validate:"required"`
	Jobs     []Job
}
