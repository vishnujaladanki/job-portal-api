package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"
)

func (s *Conn) CreateJ(ctx context.Context, nj models.NewJob, cId int) (models.Job, error) {

	job := models.Job{
		Title:       nj.Title,
		Description: nj.Description,
		CompanyID:   uint(cId),
	}

	tx := s.db.WithContext(ctx).Create(&job)

	if tx.Error != nil {
		return models.Job{}, errors.New("creation of job failed")
	}

	return job, nil
}
func (s *Conn) ViewJobs() ([]models.Job, error) {
	var jobs []models.Job

	err := s.db.Find(&jobs).Error

	if err != nil {
		return []models.Job{}, err
	}

	return jobs, nil
}
func (s *Conn) GetJobById(jId int) (models.Job, error) {
	var job models.Job
	tx := s.db.Where("ID = ?", jId)
	err := tx.Find(&job).Error
	if err != nil {
		return models.Job{}, errors.New("company not found")
	}
	return job, nil
}
func (s *Conn) ViewJobById(cId int) ([]models.Job, error) {
	var jobs []models.Job

	tx := s.db.Where("company_id =?", cId)
	err := tx.Find(&jobs).Error

	if err != nil {
		return []models.Job{}, errors.New("no jobs for that company")
	}

	return jobs, nil
}
