package repository

import (
	"context"
	"errors"
	"fmt"
	"job-portal/internal/models"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateUser is a method that creates a new user record in the database.
func (s *Conn) CreateU(ctx context.Context, nu models.NewUser) (models.User, error) {

	// We hash the user's password for storage in the database.
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}

	// We prepare the User record.
	u := models.User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashedPass),
		//CompanyID:    nu.CompanyID,
	}

	// We attempt to create the new User record in the database.
	err = s.db.Create(&u).Error
	if err != nil {
		return models.User{}, err
	}

	// Successfully created the record, return the user.
	return u, nil
}

// Authenticate is a method that checks a user's provided email and password against the database.
func (s *Conn) AuthenticateUser(ctx context.Context, email, password string) (jwt.RegisteredClaims,
	error) {

	// We attempt to find the User record where the email
	// matches the provided email.
	var u models.User
	tx := s.db.Where("email = ?", email).First(&u)
	if tx.Error != nil {
		return jwt.RegisteredClaims{}, tx.Error
	}

	// We check if the provided password matches the hashed password in the database.
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}

	// Successful authentication! Generate JWT claims.
	c := jwt.RegisteredClaims{
		Issuer:    "job-portal-api",
		Subject:   strconv.FormatUint(uint64(u.ID), 10),
		Audience:  jwt.ClaimStrings{"students"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	// And return those claims.
	return c, nil
}
func (s *Conn) CheckUserEmail(email string) (bool, error) {
	// Query the database for a user with the specified email
	var user models.User
	result := s.db.Where("email = ?", email).First(&user)

	if result.Error == nil {
		// If no error occurred, the user with the specified email exists
		return true, nil
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// If the error is ErrRecordNotFound, the user with the specified email does not exist
		return false, result.Error
	}

	// If there is an error other than ErrRecordNotFound, return the error
	return false, result.Error
}

func (s *Conn) UpdateUserPassword(np models.Reset) (bool, error) {
	var user models.User
	result := s.db.Where("email = ?", np.Email).First(&user)

	if result.Error != nil {
		return false, result.Error
	}

	// Hash the new password before updating it in the database
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(np.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return false, err
	}

	// Update the user's password in the database
	user.PasswordHash = string(hashedPassword)
	if err := s.db.Save(&user).Error; err != nil {
		return false, err
	}

	return true, nil
}
