package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateJob(ctx context.Context, nj models.NewJob, cId int) (models.Job, error) {
	job, err := r.rp.CreateJ(ctx, nj, cId)
	if err != nil {
		return models.Job{}, err
	}
	return job, nil
}

func (r NewService) ViewJob(ctx context.Context) ([]models.Job, error) {
	jobs, err := r.rp.ViewJobs()
	if err != nil {
		return []models.Job{}, err
	}
	return jobs, nil
}

func (r NewService) GetJobInfoByID(ctx context.Context, jId int) (models.Job, error) {
	job, err := r.rp.GetJobById(jId)
	if err != nil {
		return models.Job{}, err
	}
	return job, nil
}

func (r NewService) ViewJobByCompanyId(ctx context.Context, cId int) ([]models.Job, error) {
	jobs, err := r.rp.ViewJobById(cId)
	if err != nil {
		return []models.Job{}, err
	}
	return jobs, nil
}

func (r NewService) ApplyJob(application models.JobApplication, jId int) (models.Applicant, error) {
	user, err := r.rp.Process(application, jId)
	if err != nil {
		return models.Applicant{}, err
	}
	return user, nil
}
