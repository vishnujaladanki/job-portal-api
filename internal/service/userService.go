package service

import (
	"context"
	"job-portal/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func (r NewService) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {
	user, err := r.rp.CreateU(ctx, nu)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (r NewService) Authenticate(ctx context.Context, email string, password string) (jwt.RegisteredClaims, error) {
	c, err := r.rp.AuthenticateUser(ctx, email, password)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	return c, nil
}
