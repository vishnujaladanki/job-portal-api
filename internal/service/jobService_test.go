package service

import (
	"context"
	"errors"
	"job-portal/internal/models"
	"job-portal/internal/repository"
	"reflect"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestNewService_CreateJob(t *testing.T) {
	type args struct {
		ctx context.Context
		nj  models.NewJob
		cId int
	}
	tests := []struct {
		name string
		//r                NewService
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "error in creating job",
			args: args{
				ctx: context.Background(),
				nj: models.NewJob{
					Title:       "software developer",
					Description: "develop mobile applications",
					CompanyID:   24,
				},
			},
			want: models.Job{},
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("error in creating job")
			},
			wantErr: true,
		},

		{
			name: "success",
			args: args{
				ctx: context.Background(),

				nj: models.NewJob{
					Title:       "software developer",
					Description: "develop mobile applications",
					CompanyID:   24,
				},
			},

			want: models.Job{
				Title:       "software developer",
				Description: "develop mobile applications",
				CompanyID:   24,
			},

			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Title:       "software developer",
					Description: "develop mobile applications",
					CompanyID:   24,
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
				mockRepo.EXPECT().CreateJ(tt.args.ctx, tt.args.nj, tt.args.cId).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s := NewServiceStore(mockRepo)
			got, err := s.CreateJob(tt.args.ctx, tt.args.nj, tt.args.cId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.CreateJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.CreateJob() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_ViewJob(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name             string
		r                NewService
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "error in db",
			want: []models.Job{},
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{}, errors.New("error in accessing the db")
			},
			wantErr: true,
		},
		{
			name: "success",
			want: []models.Job{
				{
					Title:       "software developer",
					Description: "develop mobile applications",
					CompanyID:   24,
				},
			},
			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Title:       "software developer",
						Description: "develop mobile applications",
						CompanyID:   24,
					},
				}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mc := gomock.NewController(t)
			mockRepo := repository.NewMockRepository(mc)
			if tt.mockRepoResponse != nil {
				mockRepo.EXPECT().ViewJobs().Return(tt.mockRepoResponse()).AnyTimes()

			}

			s := NewServiceStore(mockRepo)

			got, err := s.ViewJob(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.ViewJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.ViewJob() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestNewService_GetJobInfoByID(t *testing.T) {
	type args struct {
		ctx context.Context
		jId int
	}
	tests := []struct {
		name             string
		r                NewService
		args             args
		want             models.Job
		wantErr          bool
		mockRepoResponse func() (models.Job, error)
	}{
		{
			name: "error from db",
			want: models.Job{},
			args: args{
				jId: 12,
			},
			wantErr: true,
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("test error")
			},
		},
		{
			name: "success",
			args: args{
				jId: 12,
			},
			want: models.Job{
				Title:       "software developer",
				Description: "mobile application developemt",
			},
			mockRepoResponse: func() (models.Job, error) {
				return models.Job{
					Title:       "software developer",
					Description: "mobile application developemt",
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
				mockRepo.EXPECT().GetJobById(tt.args.jId).Return(tt.mockRepoResponse()).AnyTimes()
			}

			s := NewServiceStore(mockRepo)
			got, err := s.GetJobInfoByID(tt.args.ctx, tt.args.jId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.GetJobInfoByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.GetJobInfoByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewService_ViewJobByCompanyId(t *testing.T) {
	type args struct {
		ctx context.Context
		cId int
	}
	tests := []struct {
		name             string
		r                NewService
		args             args
		want             []models.Job
		wantErr          bool
		mockRepoResponse func() ([]models.Job, error)
	}{
		{
			name: "error in db",
			args: args{
				cId: 12,
			},
			want: []models.Job{},

			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{}, errors.New("error in accesing data from db")
			},

			wantErr: true,
		},

		{
			name: "success",
			args: args{
				cId: 12,
			},
			want: []models.Job{
				{
					Title:       "software developer",
					Description: "develop mobile apps",
					CompanyID:   12,
				},
			},

			mockRepoResponse: func() ([]models.Job, error) {
				return []models.Job{
					{
						Title:       "software developer",
						Description: "develop mobile apps",
						CompanyID:   12,
					},
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
				mockRepo.EXPECT().ViewJobById(tt.args.cId).Return(tt.mockRepoResponse()).AnyTimes()
			}
			s := NewServiceStore(mockRepo)
			got, err := s.ViewJobByCompanyId(tt.args.ctx, tt.args.cId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.ViewJobByCompanyId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.ViewJobByCompanyId() = %v, want %v", got, tt.want)
			}
		})
	}
}
