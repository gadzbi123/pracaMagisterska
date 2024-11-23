package utils

import (
	"errors"
	"strconv"
	"testing"
)

func TestParseBuffer(t *testing.T) {
	scenarios := []struct {
		name          string
		input         string
		expected      int
		expectedError error
	}{
		{
			name:          "regular int",
			input:         "123",
			expected:      123,
			expectedError: nil,
		},
		{
			name:          "regular int with B",
			input:         "123B",
			expected:      123,
			expectedError: nil,
		},
		{
			name:          "10kb lowercase",
			input:         "10kb",
			expected:      int(10 * KB),
			expectedError: nil,
		},
		{
			name:          "10KB uppercase",
			input:         "10KB",
			expected:      int(10 * KB),
			expectedError: nil,
		},
		{
			name:          "512MB value",
			input:         "512MB",
			expected:      int(512 * MB),
			expectedError: nil,
		},
		{
			name:          "2048GB value",
			input:         "2048GB",
			expected:      int(2048 * GB),
			expectedError: nil,
		},
		{
			name:          "overflow case",
			input:         "9223372036854775807KB",
			expected:      0,
			expectedError: ErrOverflow,
		},
		{
			name:          "invalid format",
			input:         "abc",
			expected:      0,
			expectedError: strconv.ErrSyntax,
		},
		{
			name:          "empty input",
			input:         "",
			expected:      0,
			expectedError: strconv.ErrSyntax,
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			t.Parallel()
			result, err := ParseBufferSize(scenario.input)
			if result != scenario.expected || !errors.Is(err, scenario.expectedError) {
				t.Errorf("Test %s failed, result=%v, expected=%v, err=%v, expectedErr=%v",
					scenario.name, result, scenario.expected, err, scenario.expectedError)
			}
		})
	}

}
