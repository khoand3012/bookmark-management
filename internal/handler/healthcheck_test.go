package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/khoand3012/bookmark-management/internal/service"
	service_mocks "github.com/khoand3012/bookmark-management/internal/service/mocks"
	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		setupMockSvc     func(ctx context.Context) *service_mocks.HealthCheck
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",
			setupMockSvc: func(ctx context.Context) *service_mocks.HealthCheck {
				mockSvc := service_mocks.NewHealthCheck(t)
				mockSvc.On("CheckHealth").Return(&service.HealthCheckResult{
					Message:     "OK",
					ServiceName: "bookmark-service",
					InstanceId:  "random-string-1234",
				})
				return mockSvc
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"message": "OK", "service_name": "bookmark-service", "instance_id": "random-string-1234"}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			ginCtx, _ := gin.CreateTestContext(rec)

			mockSvc := tc.setupMockSvc(ginCtx)
			healthCheckHandler := NewHealthCheck(mockSvc)
			healthCheckHandler.CheckHealth(ginCtx)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
