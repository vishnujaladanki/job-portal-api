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

func TestNewService_ApplyJob(t *testing.T) {
	type args struct {
		application []models.JobApplication
		jId         int
	}
	tests := []struct {
		name string
		//r       NewService
		args         args
		want         []models.Applicant
		wantErr      bool
		mockResponse func() (models.Job, error)
	}{
		{
			name: "sucess case function",
			args: args{
				application: []models.JobApplication{
					{
						Name:            "vishnu",
						Email:           "vishnu@gmail.com",
						Age:             24,
						NoticePeriod:    3,
						Expect_salary:   500000,
						JobLocations:    []uint{1, 2},
						TechnologyStack: []uint{1, 2},
						WorkModes:       []uint{1, 2},
						Experience:      3,
						Qualifications:  []uint{1, 2},
						WorkShifts:      []uint{1, 2},
						JobTypes:        []uint{1, 2},
					},
					{
						Name:            "krishna",
						Email:           "krishna@gmail.com",
						Age:             24,
						NoticePeriod:    3,
						Expect_salary:   500000,
						JobLocations:    []uint{1, 2},
						TechnologyStack: []uint{1, 2},
						WorkModes:       []uint{1, 2},
						Experience:      3,
						Qualifications:  []uint{1, 2},
						WorkShifts:      []uint{1, 2},
						JobTypes:        []uint{1, 2},
					},
					{
						Name:            "vikram",
						Email:           "vikram@gmail.com",
						Age:             24,
						NoticePeriod:    3,
						Expect_salary:   500000,
						JobLocations:    []uint{3},
						TechnologyStack: []uint{3},
						WorkModes:       []uint{3},
						Experience:      3,
						Qualifications:  []uint{3},
						WorkShifts:      []uint{3},
						JobTypes:        []uint{3},
					},
				},
				jId: 1,
			},
			want:    []models.Applicant{{Name: "vishnu", Email: "vishnu@gmail.com", Age: 24}, {Name: "krishna", Email: "krishna@gmail.com", Age: 24}},
			wantErr: false,
			mockResponse: func() (models.Job, error) {
				return models.Job{
					ID:              1,
					Title:           "Java Developper",
					Description:     "train hire and deploy",
					CompanyID:       3,
					Min_NP:          1,
					Max_NP:          3,
					Budget:          500000,
					JobLocations:    []models.Location{{ID: 1, Name: "banglore"}, {ID: 2, Name: "hyderabad"}},
					TechnologyStack: []models.Technology{{ID: 1, Name: "java"}, {ID: 2, Name: "sql"}},
					WorkModes:       []models.WorkMode{{ID: 1, Name: "remote"}, {ID: 2, Name: "work from office"}},
					MinExp:          1,
					MaxExp:          3,
					Qualifications:  []models.Qualification{{ID: 1, Name: "B.Tech"}, {ID: 2, Name: "M.Tech"}},
					Shifts:          []models.Shift{{ID: 1, Name: "day"}, {ID: 2, Name: "night"}},
					JobTypes:        []models.JobType{{ID: 1, Name: "part time"}, {ID: 2, Name: "contract"}},
				}, nil
			},
		},

		{
			name: "error in db",
			args: args{
				application: []models.JobApplication{
					{
						Name:            "vishnu",
						Email:           "vishnu@gmail.com",
						Age:             24,
						NoticePeriod:    3,
						Expect_salary:   500000,
						JobLocations:    []uint{1, 2},
						TechnologyStack: []uint{1, 2},
						WorkModes:       []uint{1, 2},
						Experience:      3,
						Qualifications:  []uint{1, 2},
						WorkShifts:      []uint{1, 2},
						JobTypes:        []uint{1, 2},
					},
					{
						Name:            "vikram",
						Email:           "vikram@gmail.com",
						Age:             24,
						NoticePeriod:    3,
						Expect_salary:   500000,
						JobLocations:    []uint{3},
						TechnologyStack: []uint{3},
						WorkModes:       []uint{3},
						Experience:      3,
						Qualifications:  []uint{3},
						WorkShifts:      []uint{3},
						JobTypes:        []uint{3},
					},
				},
				jId: 1,
			},
			want:    nil,
			wantErr: true,
			mockResponse: func() (models.Job, error) {
				return models.Job{}, errors.New("")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mc := gomock.NewController(t)
			ms := repository.NewMockRepository(mc)
			if tt.mockResponse != nil {
				ms.EXPECT().Process(gomock.Any()).Return(tt.mockResponse()).AnyTimes()
			}
			r := NewServiceStore(ms)

			got, err := r.ApplyJob(tt.args.application, tt.args.jId)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewService.ApplyJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService.ApplyJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
