package antrian

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		param    string
		expected bool
	}{
		{
			param:    "0",
			expected: false,
		},
		{
			param:    "1",
			expected: true,
		},
		{
			param:    "5",
			expected: true,
		},
		{
			param:    "6",
			expected: false,
		},
	}

	for i, v := range tests {
		isValid := Validate(v.param)
		if isValid != v.expected {
			t.Errorf("fail when validated test cases on case no %d", i+1)
		}
	}
}
