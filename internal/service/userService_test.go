package service

import (
	"context"
	"errors"
	"job-portal/internal/models"
	"job-portal/internal/repository"
	"reflect"
	"testing"

	"github.com/golang-jwt/jwt/v5"
	gomock "go.uber.org/mock/gomock"
)

func TestNewService_CreateUser(t *testing.T) {
	type args struct {
		ctx context.Context
		nu  models.NewUser
	}
	tests := []struct {
		name string
		//r                NewService
		args             args
		want             models.User
		wantErr          bool
		mockRepoResponse func() (models.User, error)
	}{
		{
			name: "error in creating",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "vishnu",
					Email:    "vishnu@gmail.com",
					Password: "1234",
				},
			},
			want: models.User{},
			mockRepoResponse: func() (models.User, error) {
				return models.User{}, errors.New("error in creating user in db")
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx: context.Background(),
				nu: models.NewUser{
					Name:     "vishnu",
					Email:    "vishnu@gmail.com",
					Password: "1234",
				},
			},
			want: models.User{
				Name:  "vishnu",
				Email: "vishnu@gmail.com",
			},
			mockRepoResponse: func() (models.User, error) {
				return models.User{
					Name:  "vishnu",
					Email: "vishnu@gmail.com",
				}, nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().CreateU(tt.args.ctx, tt.args.nu).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s := NewServiceStore(mockRepo)
			got, err := s.CreateUser(tt.args.ctx, tt.args.nu)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.CreateUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_Authenticate(t *testing.T) {
	type args struct {
		ctx      context.Context
		email    string
		password string
	}
	tests := []struct {
		name string
		//r       NewService
		args             args
		want             jwt.RegisteredClaims
		wantErr          bool
		mockRepoResponse func() (jwt.RegisteredClaims, error)
	}{
		{
			name: "error in authentication",
			args: args{
				ctx:      context.Background(),
				email:    "vishnu@gmail.com",
				password: "1234",
			},
			want: jwt.RegisteredClaims{},
			mockRepoResponse: func() (jwt.RegisteredClaims, error) {
				return jwt.RegisteredClaims{}, errors.New("error while authenticating")
			},
			wantErr: true,
		},
		{
			name: "success",
			args: args{
				ctx:      context.Background(),
				email:    "vishnu@gmail.com",
				password: "1234",
			},
			want: jwt.RegisteredClaims{
				ID:        "123",
				Issuer:    "user",
				Subject:   "login",
				ExpiresAt: &jwt.NumericDate{},
				NotBefore: &jwt.NumericDate{},
				IssuedAt:  &jwt.NumericDate{},
			},

			mockRepoResponse: func() (jwt.RegisteredClaims, error) {
				return jwt.RegisteredClaims{
					ID:        "123",
					Issuer:    "user",
					Subject:   "login",
					ExpiresAt: &jwt.NumericDate{},
					NotBefore: &jwt.NumericDate{},
					IssuedAt:  &jwt.NumericDate{},
				}, nil
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			ms := repository.NewMockRepository(mc)
			if tt.mockRepoResponse != nil {
				ms.EXPECT().AuthenticateUser(tt.args.ctx, tt.args.email, tt.args.password).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s := NewServiceStore(ms)
			got, err := s.Authenticate(tt.args.ctx, tt.args.email, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.Authenticate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.Authenticate() = %v, want %v", got, tt.want)
			}
		})
	}
}
