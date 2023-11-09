package service

import (
	"context"
	"errors"
	"job-portal/internal/models"

	"github.com/rs/zerolog/log"
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
	user := models.Applicant{
		Name:  application.Name,
		Email: application.Email,
		Age:   application.Age,
	}
	var count int
	job, err := r.rp.Process(jId)
	if err != nil {
		return models.Applicant{}, err
	}
	err = errors.New("")
	if application.Expect_salary <= job.Budget {
		log.Info().Str("Budget", "true").Send()
		count++
	} else {
		log.Info().Str("Budget", "false").Send()
		return models.Applicant{}, err
	}
	if application.NoticePeriod >= job.Min_NP && application.NoticePeriod <= job.Max_NP {
		log.Info().Str("Min_NP", "true").Send()
		count++
	} else {
		log.Info().Str("Min_NP", "false").Send()
		return models.Applicant{}, err
	}

	if application.Experience >= job.MinExp && application.Experience <= job.MaxExp {
		log.Info().Str("MinExp", "true").Send()
		count++
	} else {
		log.Info().Str("MinExp", "false").Send()
		return models.Applicant{}, err
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

	if count >= 4 {
		return user, nil
	}

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
