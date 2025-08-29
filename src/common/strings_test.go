package common

import (
	"regexp"
	"testing"
)

// Mock config for testing
type mockPasswordConfig struct {
	MinLength        int
	IncludeChars     bool
	IncludeDigits    bool
	IncludeLowercase bool
	IncludeUppercase bool
}
type mockOtpConfig struct {
	Digits int
}
type mockConfig struct {
	Password mockPasswordConfig
	Otp      mockOtpConfig
}

var testConfig = &mockConfig{
	Password: mockPasswordConfig{
		MinLength:        8,
		IncludeChars:     true,
		IncludeDigits:    true,
		IncludeLowercase: true,
		IncludeUppercase: true,
	},
	Otp: mockOtpConfig{
		Digits: 6,
	},
}

// Patch config.GetConfig for tests
func init() {
	configGetConfig = func() *mockConfig {
		return testConfig
	}
}

// Patch point for config.GetConfig
var configGetConfig = func() *mockConfig { return testConfig }

// Patch config.GetConfig in tested code
func getConfig() *mockConfig {
	return configGetConfig()
}

// --- Test Functions ---

func TestCheckPassword(t *testing.T) {
	testConfig.Password = mockPasswordConfig{
		MinLength:        8,
		IncludeChars:     true,
		IncludeDigits:    true,
		IncludeLowercase: true,
		IncludeUppercase: true,
	}
	tests := []struct {
		password string
		want     bool
	}{
		{"Abcdef12!", true},
		{"abcdef12", false}, // no uppercase
		{"ABCDEF12", false}, // no lowercase
		{"Abcdefgh", false}, // no digits
		{"Abc12", false},    // too short
		{"12345678", false}, // no letters
	}
	for _, tt := range tests {
		got := CheckPassword(tt.password)
		if got != tt.want {
			t.Errorf("CheckPassword(%q) = %v, want %v", tt.password, got, tt.want)
		}
	}
}

func TestGeneratePassword(t *testing.T) {
	testConfig.Password = mockPasswordConfig{
		MinLength:        8,
		IncludeChars:     true,
		IncludeDigits:    true,
		IncludeLowercase: true,
		IncludeUppercase: true,
	}
	pass := GeneratePassword()
	if len(pass) < testConfig.Password.MinLength {
		t.Errorf("Generated password too short: %s", pass)
	}
	if !HasLetter(pass) {
		t.Errorf("Generated password missing letter: %s", pass)
	}
	if !HasDigits(pass) {
		t.Errorf("Generated password missing digit: %s", pass)
	}
	if !HasLower(pass) {
		t.Errorf("Generated password missing lowercase: %s", pass)
	}
	if !HasUpper(pass) {
		t.Errorf("Generated password missing uppercase: %s", pass)
	}
}

func TestGenerateOtp(t *testing.T) {
	testConfig.Otp.Digits = 6
	otp := GenerateOtp()
	if len(otp) != testConfig.Otp.Digits {
		t.Errorf("GenerateOtp() = %s, want length %d", otp, testConfig.Otp.Digits)
	}
	if !regexp.MustCompile(`^\d+$`).MatchString(otp) {
		t.Errorf("GenerateOtp() = %s, want only digits", otp)
	}
}

func TestHasUpper(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"abc", false},
		{"ABC", true},
		{"aBc", true},
		{"123", false},
	}
	for _, tt := range tests {
		got := HasUpper(tt.in)
		if got != tt.want {
			t.Errorf("HasUpper(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestHasLower(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"abc", true},
		{"ABC", false},
		{"aBc", true},
		{"123", false},
	}
	for _, tt := range tests {
		got := HasLower(tt.in)
		if got != tt.want {
			t.Errorf("HasLower(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestHasLetter(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"abc", true},
		{"123", false},
		{"a1b2", true},
		{"!@#", false},
	}
	for _, tt := range tests {
		got := HasLetter(tt.in)
		if got != tt.want {
			t.Errorf("HasLetter(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestHasDigits(t *testing.T) {
	tests := []struct {
		in   string
		want bool
	}{
		{"abc", false},
		{"123", true},
		{"a1b2", true},
		{"!@#", false},
	}
	for _, tt := range tests {
		got := HasDigits(tt.in)
		if got != tt.want {
			t.Errorf("HasDigits(%q) = %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestToSnakeCase(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"CountryId", "country_id"},
		{"UserID", "user_id"},
		{"SomeLongFieldName", "some_long_field_name"},
		{"simple", "simple"},
		{"HTTPServer", "http_server"},
	}
	for _, tt := range tests {
		got := ToSnakeCase(tt.in)
		if got != tt.want {
			t.Errorf("ToSnakeCase(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}
