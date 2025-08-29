package common

import (
	"testing"
)

func TestIranianMobileNumberValidate(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		// Valid Iranian mobile numbers
		{"09121234567", true},
		{"09351234567", true},
		{"09901234567", true},
		{"09221234567", true},
		{"09111234567", true},
		// Invalid Iranian mobile numbers
		{"08121234567", false},  // wrong prefix
		{"0912123456", false},   // too short
		{"091212345678", false}, // too long
		{"0912123456a", false},  // contains letter
		{"09531234567", false},  // invalid operator code
		{"", false},             // empty string
	}

	for _, tt := range tests {
		result := IranianMobileNumberValidate(tt.input)
		if result != tt.expected {
			t.Errorf("IranianMobileNumberValidate(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
