package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "birthday already passed this year",
			dob:      now.AddDate(-20, -1, 0),
			expected: 20,
		},
		{
			name:     "birthday not yet passed this year",
			dob:      now.AddDate(-20, 1, 0),
			expected: 19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := calculateAge(tt.dob)

			if age != tt.expected {
				t.Errorf("expected %d, got %d", tt.expected, age)
			}
		})
	}
}
