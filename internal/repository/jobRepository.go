package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"
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

func (s *Conn) Process(jId int) (models.Job, error) {
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
		return models.Job{}, err
	}
	return job, nil
}
