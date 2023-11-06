package service

import (
	"context"
	"job-portal/internal/models"
)

func (r NewService) CreateCompany(ctx context.Context, ni models.NewCompany) (models.Company, error) {
	c, err := r.rp.CreateC(ctx, ni)
	if err != nil {
		return models.Company{}, err
	}
	return c, nil
}

func (r NewService) ViewCompany(ctx context.Context) ([]models.Company, error) {
	c, err := r.rp.ViewCompanies()
	if err != nil {
		return []models.Company{}, err
	}
	return c, nil
}

func (r NewService) GetCompanyInfoByID(ctx context.Context, uid int) (models.Company, error) {
	c, err := r.rp.GetCompanyByID(uid)
	if err != nil {
		return models.Company{}, err
	}
	return c, nil
}
