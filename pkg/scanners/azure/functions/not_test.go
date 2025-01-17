package functions

import "testing"

func Test_Not(t *testing.T) {
	tests := []struct {
		name     string
		args     []interface{}
		expected bool
	}{
		{
			name:     "Not with true",
			args:     []interface{}{true},
			expected: false,
		},
		{
			name:     "Not with false",
			args:     []interface{}{false},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Not(tt.args...)
			if got != tt.expected {
				t.Errorf("Not() = %v, want %v", got, tt.expected)
			}
		})
	}
}
