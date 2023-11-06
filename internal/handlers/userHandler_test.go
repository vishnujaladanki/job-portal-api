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

func Test_handler_UserRegister(t *testing.T) {
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
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		{
			name: "error while adding user",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{"name": "vishnu", "email":"vishnu@gmail.com", "password":"123456"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateUser(c.Request.Context(), gomock.Any()).Return(models.User{}, errors.New("error in adding job")).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   `{"error":"user signup failed"}`,
		},
		{
			name: "sucessfully adding user",
			setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				requestBody := []byte(`{"name": "vishnu", "email":"vishnu@gmail.com", "password":"123456"}`)
				httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
				ctx := httpReq.Context()
				ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
				httpReq = httpReq.WithContext(ctx)
				c.Request = httpReq

				mc := gomock.NewController(t)
				ms := service.NewMockService(mc)
				ms.EXPECT().CreateUser(c.Request.Context(), gomock.Any()).Return(models.User{}, nil).AnyTimes()

				return c, rr, ms
			},
			expectedStatusCode: http.StatusOK,
			expectedResponse:   `{"ID":0,"CreatedAt":"0001-01-01T00:00:00Z","UpdatedAt":"0001-01-01T00:00:00Z","DeletedAt":null,"UserId":0,"name":"","email":""}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.UserRegister(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}

func Test_handler_UserLogin(t *testing.T) {
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
				c.Request = httpReq

				return c, rr, nil
			},
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   `{"error":"Bad Request"}`,
		},
		// {
		// 	name: "error while athenticating user",
		// 	setup: func() (*gin.Context, *httptest.ResponseRecorder, service.Service) {
		// 		rr := httptest.NewRecorder()
		// 		c, _ := gin.CreateTestContext(rr)
		// 		requestBody := []byte(`{"email":"vishnu@gmail.com", "password":"123456"}`)
		// 		httpReq, _ := http.NewRequest(http.MethodGet, "http://google.com:8080", bytes.NewBuffer(requestBody))
		// 		ctx := httpReq.Context()
		// 		ctx = context.WithValue(ctx, middleware.TraceIdKey, "693")
		// 		httpReq = httpReq.WithContext(ctx)
		// 		c.Request = httpReq

		// 		mc := gomock.NewController(t)
		// 		ms := service.NewMockService(mc)
		// 		ms.EXPECT().Authenticate(c.Request.Context(), gomock.Any(), gomock.Any())(jwt.RegisteredClaims{}, errors.New("")).AnyTimes()

		// 		return c, rr, nil
		// 	},
		// 	expectedStatusCode: http.StatusInternalServerError,
		// 	expectedResponse:   `{"error":"user signup failed"}`,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			c, rr, ms := tt.setup()
			h := &handler{
				s: ms,
			}
			h.UserLogin(c)
			assert.Equal(t, tt.expectedStatusCode, rr.Code)
			assert.Equal(t, tt.expectedResponse, rr.Body.String())
		})
	}
}
