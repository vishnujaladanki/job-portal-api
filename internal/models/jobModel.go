package models

import "gorm.io/gorm"

type Location struct {
	gorm.Model
	ID   uint
	Name string
}

type Technology struct {
	gorm.Model
	ID   uint
	Name string
}

type WorkMode struct {
	gorm.Model
	ID   uint
	Name string
}

type Qualification struct {
	gorm.Model
	ID   uint
	Name string
}

type Shift struct {
	gorm.Model
	ID   uint
	Name string
}

type JobType struct {
	gorm.Model
	ID   uint
	Name string
}

type Job struct {
	gorm.Model
	ID              uint `gorm:"primaryKey;autoIncrement"`
	Title           string
	Description     string
	CompanyID       uint
	Min_NP          int
	Max_NP          int
	Budget          int
	JobLocations    []Location   `gorm:"many2many:location_jobs"`
	TechnologyStack []Technology `gorm:"many2many:technology_jobs"`
	WorkModes       []WorkMode   `gorm:"many2many:workmode_jobs"`
	MinExp          int
	MaxExp          int
	Qualifications  []Qualification `gorm:"many2many:qualification_jobs"`
	Shifts          []Shift         `gorm:"many2many:shift_jobs"`
	JobTypes        []JobType       `gorm:"many2many:jobtype_jobs"`
}

type NewJob struct {
	Title           string `json:"title" validate:"required"`
	Description     string `json:"description" validate:"required"`
	CompanyID       uint
	Min_NP          int    `json:"min_np" validate:"required"`
	Max_NP          int    `json:"max_np" validate:"required"`
	Budget          int    `json:"budget" validate:"required"`
	JobLocations    []uint `json:"job_location" validate:"required"`
	TechnologyStack []uint `json:"technology_stack" validate:"required"`
	WorkModes       []uint `json:"work_mode" validate:"required"`
	MinExp          int    `json:"min_exp" validate:"required"`
	MaxExp          int    `json:"max_exp" validate:"required"`
	Qualifications  []uint `json:"qualification" validate:"required"`
	WorkShifts      []uint `json:"work_shift" validate:"required"`
	JobTypes        []uint `json:"job_type" validate:"required"`
}

type JobApplication struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Age             int    `json:"age" validate:"required"`
	JobId           int    `json:"job_Id" validate:"required"`
	NoticePeriod    int    `json:"notice_period" validate:"required"`
	Expect_salary   int    `json:"expect_salary" validate:"required"`
	JobLocations    []uint `json:"job_location" validate:"required"`
	TechnologyStack []uint `json:"technology_stack" validate:"required"`
	WorkModes       []uint `json:"work_mode" validate:"required"`
	Experience      int    `json:"experience" validate:"required"`
	Qualifications  []uint `json:"qualification" validate:"required"`
	WorkShifts      []uint `json:"work_shift" validate:"required"`
	JobTypes        []uint `json:"job_type" validate:"required"`
}

type Applicant struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
	Age   int    `json:"age" validate:"required"`
}
