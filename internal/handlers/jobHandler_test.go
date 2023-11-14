package handlers

import (
	"bytes"
	"context"
	"errors"
	"job-portal/internal/middleware"
	"job-portal/internal/models"
	"job-portal/internal/service"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"
	"gopkg.in/go-playground/assert.v1"
)

func Test_handler_ViewJobByCompany(t *testing.T) {

	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "Invalid companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "vishnu"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while fectching job details by companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJobByCompanyId(c.Request.Context(), gomock.Any()).Return([]models.Job{}, errors.New("mock service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"mock service error"}`,
		},
		{
			name: "sucess while fectching job details by companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJobByCompanyId(c.Request.Context(), gomock.Any()).Return([]models.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ViewJobByCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ViewJobById(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "Invalid jobId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "vishnu"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while fectching job details by jobId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().GetJobInfoByID(c.Request.Context(), gomock.Any()).Return(models.Job{}, errors.New("mock service error")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"mock service error"}`,
		},
		{
			name: "sucess while fectching job details by jobId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "693"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().GetJobInfoByID(c.Request.Context(), gomock.Any()).Return(models.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"ID":0,"Title":"","Description":"","CompanyID":0,"Min_NP":0,"Max_NP":0,"Budget":0,"JobLocations":null,"TechnologyStack":null,"WorkModes":null,"MinExp":0,"MaxExp":0,"Qualifications":null,"Shifts":null,"JobTypes":null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ViewJobById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_AddJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{

		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "Invalid CompanyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "vishnu"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "invalid request body",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := "invalid string request body"
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", strings.NewReader(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "checking validator function",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{"key": "value"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while adding job",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{
					"title": "java developer",
					"Description": "hired by train hire and deploy program",
					"min_np": 1,
					"max_np": 5,
					"budget": 500000,
					"job_location": [3],
					"technology_stack": [3,4], 
					"work_mode": [3],
					"min_exp": 3,
					"max_exp": 8,
					"qualification": [3,4],
					"work_shift": [3],
					"job_type": [3]
				  }`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateJob(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.Job{}, errors.New("error in adding job")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"error in adding job"}`,
		},
		{
			name: "sucessfully adding job",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{
					"title": "java developer",
					"Description": "hired by train hire and deploy program",
					"min_np": 1,
					"max_np": 5,
					"budget": 500000,
					"job_location": [3],
					"technology_stack": [3,4], 
					"work_mode": [3],
					"min_exp": 3,
					"max_exp": 8,
					"qualification": [3,4],
					"work_shift": [3],
					"job_type": [3]
				  }`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateJob(c.Request.Context(), gomock.Any(), gomock.Any()).Return(models.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"ID":0,"Title":"","Description":"","CompanyID":0,"Min_NP":0,"Max_NP":0,"Budget":0,"JobLocations":null,"TechnologyStack":null,"WorkModes":null,"MinExp":0,"MaxExp":0,"Qualifications":null,"Shifts":null,"JobTypes":null}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.AddJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ViewJobs(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "error while fectching jobs",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJob(c.Request.Context()).Return([]models.Job{}, errors.New("jobs not found")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"jobs not found"}`,
		},
		{
			name: "sucess while fectching jobs",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq := httptest.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ViewJob(c.Request.Context()).Return([]models.Job{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ViewJobs(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ApplyForJob(t *testing.T) {
	tests := []struct {
		name               string
		setup              func() (*gin.Context, *httptest.ResponseRecorder, service.Service)
		expectedStatusCode int
		expectedResponse   string
	}{
		{
			name: "missing trace id",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com", nil)
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"Internal Server Error"}`,
		},
		{
			name: "Invalid jobId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "vishnu"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "checking decode function",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := "invalid string request body"
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", strings.NewReader(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "checking validator function",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`[{
					"name": "vishnu",
					"email": "vishnu@example.com",
					"age": 30,
					"notice_period": "1",
					"expect_salary": 500000,
					"job_location": [1, 2],
					"technology_stack": [1,2],
					"work_mode": [1,2],
					"experience": 8,
					"qualification": [1,2,3],
					"work_shift": [1,2],
					"job_type": [1,2]
				  },
				  {
					"name": "vishnu",
					"email": "vishnu@example.com",
					"age": 30,
					"notice_period": 1,
					"expect_salary": 500000,
					"job_location": [1, 2],
					"technology_stack": [1,2],
					"work_mode": [1,2],
					"experience": 8,
					"qualification": [1,2,3],
					"work_shift": [1,2],
					"job_type": [1,2]
				  }]`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},

		{
			name: "error while applying job",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`[{
					"name": "vishnu",
					"email": "vishnu@example.com",
					"age": 30,
					"notice_period": 1,
					"expect_salary": 500000,
					"job_location": [1, 2],
					"technology_stack": [1,2],
					"work_mode": [1,2],
					"experience": 8,
					"qualification": [1,2,3],
					"work_shift": [1,2],
					"job_type": [1,2]
				  },
				  {
					"name": "vishnu",
					"email": "vishnu@example.com",
					"age": 30,
					"notice_period": 1,
					"expect_salary": 500000,
					"job_location": [1, 2],
					"technology_stack": [1,2],
					"work_mode": [1,2],
					"experience": 8,
					"qualification": [1,2,3],
					"work_shift": [1,2],
					"job_type": [1,2]
				  }]`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ApplyJob(gomock.Any(), gomock.Any()).Return([]models.Applicant{}, errors.New("error in adding job"))

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"error in adding job"}`,
		},
		{
			name: "sucess while applying job",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`[{
					"name": "vishnu",
					"email": "vishnu@example.com",
					"age": 30,
					"notice_period": 1,
					"expect_salary": 500000,
					"job_location": [1, 2],
					"technology_stack": [1,2],
					"work_mode": [1,2],
					"experience": 8,
					"qualification": [1,2,3],
					"work_shift": [1,2],
					"job_type": [1,2]
				  },
				  {
					"name": "vishnu",
					"email": "vishnu@example.com",
					"age": 30,
					"notice_period": 1,
					"expect_salary": 500000,
					"job_location": [1, 2],
					"technology_stack": [1,2],
					"work_mode": [1,2],
					"experience": 8,
					"qualification": [1,2,3],
					"work_shift": [1,2],
					"job_type": [1,2]
				  }]`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().ApplyJob(gomock.Any(), gomock.Any()).Return([]models.Applicant{}, nil)

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `[]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.ApplyForJob(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
