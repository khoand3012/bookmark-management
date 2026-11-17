package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		expectedLength int
		expectedErr    error
	}{
		{
			name:           "normal case - valid length",
			expectedLength: 10,
			expectedErr:    nil,
		},
		{
			name:           "normal case 2 - valid length",
			expectedLength: 15,
			expectedErr:    nil,
		},
		{
			name:           "edge case - invalid length",
			expectedLength: 1000000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			svc := NewGenPass()

			password, err := svc.GeneratePassword(tc.expectedLength)
			assert.ErrorIs(t, err, tc.expectedErr)
			assert.Equal(t, tc.expectedLength, len(password))
		})
	}
}
