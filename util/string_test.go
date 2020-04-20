package util

import "testing"

func TestIsPositiveNumber(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{
			param:    "1",
			expected: true,
		},
		{
			param:    "0",
			expected: true,
		},
		{
			param:    "a",
			expected: false,
		},
	}

	for i, v := range tests {
		isPositive, _ := IsPositiveNumber(v.param)
		if isPositive != v.expected {
			t.Errorf("fail when validated test cases on case no %d", i+1)
		}
	}
}
