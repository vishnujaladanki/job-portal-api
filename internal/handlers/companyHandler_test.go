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

func Test_handler_CreateCompany(t *testing.T) {
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
			name: "error while creating a company",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{"name": "TekSystems", "location":"Banglore"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateCompany(c.Request.Context(), gomock.Any()).Return(models.Company{}, errors.New("error in creating company")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"error in creating company"}`,
		},
		{
			name: "sucessfully adding company",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{"name": "TekSystems", "location":"Banglore"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateCompany(c.Request.Context(), gomock.Any()).Return(models.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"ID":0,"Name":"","Location":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.CreateCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_ViewCompany(t *testing.T) {
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
			name: "error while fectching companies",
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
				ms.EXPECT().ViewCompany(c.Request.Context()).Return([]models.Company{}, errors.New("companies not found")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"companies not found"}`,
		},
		{
			name: "sucessfully fetching companies",
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
				ms.EXPECT().ViewCompany(c.Request.Context()).Return([]models.Company{}, nil).AnyTimes()

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
			h.ViewCompany(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_GetCompanyById(t *testing.T) {
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
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "one"})
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while fectching company details by companyId",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", nil)
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})
				c.Request = httpReq
				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().GetCompanyInfoByID(c.Request.Context(), gomock.Any()).Return(models.Company{}, errors.New("mock while fectching company")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"mock while fectching company"}`,
		},
		{
			name: "sucess while fectching company details by companyId",
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
				ms.EXPECT().GetCompanyInfoByID(c.Request.Context(), gomock.Any()).Return(models.Company{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"ID":0,"Name":"","Location":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.GetCompanyById(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
