package repository

import (
	"context"
	"errors"
	"job-portal/internal/models"
)

func (s *Conn) CreateC(ctx context.Context, nc models.NewCompany) (models.Company, error) {

	com := models.Company{
		Name:     nc.Name,
		Location: nc.Location,
		//CompanyId: nc.CompanyID,
		Jobs: nc.Jobs,
	}
	tx := s.db.WithContext(ctx).Create(&com)

	if tx.Error != nil {

		return models.Company{}, tx.Error
	}

	return com, nil
}

func (s *Conn) ViewCompanies() ([]models.Company, error) {
	var com []models.Company
	err := s.db.Find(&com).Error

	if err != nil {
		return []models.Company{}, err
	}

	return com, nil
}

func (s *Conn) GetCompanyByID(uid int) (models.Company, error) {

	var com models.Company
	tx := s.db.Where("ID = ?", uid)
	err := tx.Find(&com).Error
	if err != nil {
		return models.Company{}, errors.New("company not found")
	}
	return com, nil

}
