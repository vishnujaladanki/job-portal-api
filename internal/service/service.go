package service

import (
	"context"
	"job-portal/internal/models"
	"job-portal/internal/repository"

	"github.com/golang-jwt/jwt/v5"
)

type NewService struct {
	rp repository.Repository
}

//go:generate mockgen -source=service.go -destination=service_mock.go -package=service
type Service interface {
	CreateUser(ctx context.Context, nu models.NewUser) (models.User, error)
	Authenticate(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error)
	CreateJob(ctx context.Context, nj models.NewJob, cId int) (models.Job, error)
	ViewJob(ctx context.Context) ([]models.Job, error)
	GetJobInfoByID(ctx context.Context, jId int) (models.Job, error)
	ViewJobByCompanyId(ctx context.Context, cId int) ([]models.Job, error)
	CreateCompany(ctx context.Context, ni models.NewCompany) (models.Company, error)
	ViewCompany(ctx context.Context) ([]models.Company, error)
	GetCompanyInfoByID(ctx context.Context, uid int) (models.Company, error)
	ApplyJob(application models.JobApplication, jId int) (models.Applicant, error)
}

func NewServiceStore(s repository.Repository) Service {
	return &NewService{rp: s}
}
