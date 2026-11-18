package integration_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/khoand3012/bookmark-management/internal/api"
	"github.com/stretchr/testify/assert"
)

const (
	testServiceName = "test-service"
	testInstanceId  = "test-instance-1"
)

func newTestEngine() api.Engine {
	cfg := &api.Config{
		ServiceName: testServiceName,
		InstanceId:  testInstanceId,
	}
	return api.NewEngine(cfg)
}

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		setupTestRequest func() (*httptest.ResponseRecorder, *http.Request)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "normal case",
			setupTestRequest: func() (*httptest.ResponseRecorder, *http.Request) {
				req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
				rec := httptest.NewRecorder()
				return rec, req
			},
			expectedStatus: http.StatusOK,
			expectedResponse: fmt.Sprintf(
				`{"message":"OK","service_name":%q,"instance_id":%q}`,
				testServiceName, testInstanceId,
			),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			engine := newTestEngine()
			rec, req := tc.setupTestRequest()
			engine.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.JSONEq(t, tc.expectedResponse, rec.Body.String())
		})
	}
}

func TestHealthCheckNotFound(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name             string
		setupTestRequest func() (*httptest.ResponseRecorder, *http.Request)
		expectedStatus   int
		expectedResponse string
	}{
		{
			name: "wrong method",
			setupTestRequest: func() (*httptest.ResponseRecorder, *http.Request) {
				req := httptest.NewRequest(http.MethodPost, "/health-check", nil)
				rec := httptest.NewRecorder()
				return rec, req
			},
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "404 page not found",
		},
		{
			name: "wrong endpoint",
			setupTestRequest: func() (*httptest.ResponseRecorder, *http.Request) {
				req := httptest.NewRequest(http.MethodGet, "/wrong_endpoint", nil)
				rec := httptest.NewRecorder()
				return rec, req
			},
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "404 page not found",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			engine := newTestEngine()
			rec, req := tc.setupTestRequest()
			engine.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectedResponse, rec.Body.String())
		})
	}
}
