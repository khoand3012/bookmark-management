package integration_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/khoand3012/bookmark-management/internal/api"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
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
				req := httptest.NewRequest("GET", "/genpass", nil)
				rec := httptest.NewRecorder()
				return rec, req
			},
			expectedStatus:   http.StatusOK,
			expectedResponse: `{"password":`,
		},
		{
			name: "wrong endpoint",
			setupTestRequest: func() (*httptest.ResponseRecorder, *http.Request) {
				req := httptest.NewRequest("GET", "/wrong_endpoint", nil)
				rec := httptest.NewRecorder()
				return rec, req
			},
			expectedStatus:   http.StatusNotFound,
			expectedResponse: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			cfg, err := api.NewConfig()
			if err != nil {
				panic(err)
			}
			testAPIEngine := api.NewEngine(cfg)
			rec, req := tc.setupTestRequest()
			testAPIEngine.ServeHTTP(rec, req)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), tc.expectedResponse)
		})
	}
}
