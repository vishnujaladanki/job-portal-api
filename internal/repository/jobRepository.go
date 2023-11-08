package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"

	"github.com/rs/zerolog/log"
)

func (s *Conn) CreateJ(ctx context.Context, nj models.NewJob, cId int) (models.Job, error) {
	// Create a new Job instance
	job := models.Job{
		Title:       nj.Title,
		Description: nj.Description,
		CompanyID:   uint(cId),
		Min_NP:      nj.Min_NP,
		Max_NP:      nj.Max_NP,
		Budget:      nj.Budget,
		MinExp:      nj.MinExp,
		MaxExp:      nj.MaxExp,
	}

	// Map foreign key IDs to related entities in the database
	for _, locationID := range nj.JobLocations {
		job.JobLocations = append(job.JobLocations, models.Location{ID: locationID})
	}

	for _, techID := range nj.TechnologyStack {
		job.TechnologyStack = append(job.TechnologyStack, models.Technology{ID: techID})
	}

	for _, modeID := range nj.WorkModes {
		job.WorkModes = append(job.WorkModes, models.WorkMode{ID: modeID})
	}

	for _, qualificationID := range nj.Qualifications {
		job.Qualifications = append(job.Qualifications, models.Qualification{ID: qualificationID})
	}

	for _, shiftID := range nj.WorkShifts {
		job.Shifts = append(job.Shifts, models.Shift{ID: shiftID})
	}

	for _, jobTypeID := range nj.JobTypes {
		job.JobTypes = append(job.JobTypes, models.JobType{ID: jobTypeID})
	}

	// Create the job and its relationships
	tx := s.db.WithContext(ctx).Create(&job)

	if tx.Error != nil {
		return models.Job{}, errors.New("creation of job failed")
	}

	return job, nil
}

func (s *Conn) ViewJobs() ([]models.Job, error) {
	var jobs []models.Job

	err := s.db.
		Preload("JobLocations").
		Preload("TechnologyStack").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Preload("JobTypes").
		Find(&jobs).Error

	if err != nil {
		return []models.Job{}, err
	}

	return jobs, nil
}

func (s *Conn) GetJobById(jId int) (models.Job, error) {
	var job models.Job
	tx := s.db.
		Preload("JobLocations").
		Preload("TechnologyStack").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Preload("JobTypes").
		Where("ID = ?", jId)
	err := tx.First(&job).Error
	if err != nil {
		return models.Job{}, errors.New("job not found")
	}
	return job, nil
}

func (s *Conn) ViewJobById(cId int) ([]models.Job, error) {
	var jobs []models.Job

	tx := s.db.
		Preload("JobLocations").
		Preload("TechnologyStack").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Preload("JobTypes").Where("company_id =?", cId)
	err := tx.Find(&jobs).Error

	if err != nil {
		return []models.Job{}, errors.New("no jobs for that company")
	}

	return jobs, nil
}

func (s *Conn) Process(application models.JobApplication, jId int) (models.Applicant, error) {
	user := models.Applicant{
		Name:  application.Name,
		Email: application.Email,
		Age:   application.Age,
	}
	var count int
	var job models.Job
	tx := s.db.
		Preload("JobLocations").
		Preload("TechnologyStack").
		Preload("WorkModes").
		Preload("Qualifications").
		Preload("Shifts").
		Preload("JobTypes").
		Where("ID = ?", jId)
	err := tx.First(&job).Error

	if err != nil {
		return models.Applicant{}, err
	}
	if application.Budget <= job.Budget {
		log.Info().Str("Budget", "true").Send()
		count++
	} else {
		log.Info().Str("Budget", "false").Send()
	}
	if application.Min_NP <= job.Min_NP {
		log.Info().Str("Min_NP", "true").Send()
		count++
	} else {
		log.Info().Str("Min_NP", "false").Send()
	}
	if application.Max_NP <= job.Max_NP {
		log.Info().Str("Max_NP", "true").Send()
		count++
	} else {
		log.Info().Str("Max_NP", "false").Send()
	}
	if application.MinExp >= job.MinExp && application.MinExp <= job.MaxExp {
		log.Info().Str("MinExp", "true").Send()
		count++
	} else {
		log.Info().Str("MinExp", "false").Send()
	}
	if application.MaxExp >= job.MinExp && application.MaxExp <= job.MaxExp {
		log.Info().Str("MaxExp", "true").Send()
		count++
	} else {
		log.Info().Str("MaxExp", "false").Send()
	}
	//comparing job criteria locations and application criteria locations
	var loc_job []uint
	var loc_app []uint
	for _, v := range job.JobLocations {
		loc_job = append(loc_job, v.ID)
	}
	loc_app = application.JobLocations
	if sliceContainsAtLeastOne(loc_job, loc_app) {
		log.Info().Str("JobLocations", "true").Send()
		count++
	} else {
		log.Info().Str("JobLocations", "false").Send()
	}

	//comparing job criteria technologystack and application criteria technologystack
	var tech_job []uint
	var tech_app []uint
	for _, v := range job.TechnologyStack {
		tech_job = append(tech_job, v.ID)
	}
	tech_app = application.TechnologyStack
	if sliceContainsAtLeastOne(tech_job, tech_app) {
		log.Info().Str("TechnologyStack", "true").Send()
		count++
	} else {
		log.Info().Str("TechnologyStack", "false").Send()
	}

	//comparing job criteria technologystack and application criteria technologystack
	var mode_job []uint
	var mode_app []uint
	for _, v := range job.WorkModes {
		mode_job = append(mode_job, v.ID)
	}
	mode_app = application.WorkModes
	if sliceContainsAtLeastOne(mode_job, mode_app) {
		log.Info().Str("WorkModes", "true").Send()
		count++
	} else {
		log.Info().Str("WorkModes", "false").Send()
	}

	//comparing job criteria qualification and application criteria qualification
	var q_job []uint
	var q_app []uint
	for _, v := range job.Qualifications {
		q_job = append(q_job, v.ID)
	}
	q_app = application.Qualifications
	if sliceContainsAtLeastOne(q_job, q_app) {
		log.Info().Str("Qualificvations", "true").Send()
		count++
	} else {
		log.Info().Str("Qualifications", "false").Send()
	}

	//comparing job criteria shifts and application criteria shifts
	var shift_job []uint
	var shift_app []uint
	for _, v := range job.Shifts {
		shift_job = append(shift_job, v.ID)
	}
	shift_app = application.WorkShifts
	if sliceContainsAtLeastOne(shift_job, shift_app) {
		log.Info().Str("Shifts", "true").Send()
		count++
	} else {
		log.Info().Str("Shifts", "false").Send()
	}

	//comparing job criteria technologystack and application criteria technologystack
	var type_job []uint
	var type_app []uint
	for _, v := range job.JobTypes {
		type_job = append(type_job, v.ID)
	}
	type_app = application.JobTypes
	if sliceContainsAtLeastOne(type_job, type_app) {
		log.Info().Str("JobTypes", "true").Send()
		count++
	} else {
		log.Info().Str("JobTypes", "false").Send()
	}

	if count >= 5 {
		return user, nil
	}
	err = errors.New("")
	return models.Applicant{}, err
}

// function to check the slices
func sliceContainsAtLeastOne(slice, subSlice []uint) bool {
	for _, v := range subSlice {
		for _, s := range slice {
			if v == s {
				return true
			}
		}
	}
	return false
}
