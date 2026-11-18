package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name                string
		expectedServiceName string
		expectedInstanceId  string
		expectedMessage     string
	}{
		{
			name:                "normal case - health check ok",
			expectedServiceName: "bookmark-service",
			expectedInstanceId:  "random-string-1234",
			expectedMessage:     "OK",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			svc := NewHealthCheck("bookmark-service", "random-string-1234")
			result := svc.CheckHealth()

			assert.Equal(t, tc.expectedMessage, result.Message)
			assert.Equal(t, tc.expectedServiceName, result.ServiceName)
			assert.Equal(t, tc.expectedInstanceId, result.InstanceId)
		})
	}
}
