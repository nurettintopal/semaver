package semaver

import (
	"fmt"
	"testing"
)

func TestParseVersion(t *testing.T) {
	tests := []struct {
		input    string
		expected Version
		err      bool
	}{
		{"1.2.3", Version{major: 1, minor: 2, patch: 3, prerelease: "", build: ""}, false},
		{"1.2.3-alpha", Version{major: 1, minor: 2, patch: 3, prerelease: "alpha", build: ""}, false},
		{"1.2.3+build123", Version{major: 1, minor: 2, patch: 3, prerelease: "", build: "build123"}, false},
		{"1.2.3-alpha+build123", Version{major: 1, minor: 2, patch: 3, prerelease: "alpha", build: "build123"}, false},
		{"1.2", Version{}, true},     // Invalid version format
		{"1.2.3.4", Version{}, true}, // Invalid version format
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			v, err := ParseVersion(tt.input)
			if (err != nil) != tt.err {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if v != tt.expected {
				t.Errorf("expected: %+v, got: %+v", tt.expected, v)
			}
		})
	}
}

func TestCompareVersions(t *testing.T) {
	tests := []struct {
		v1       string
		operator string
		v2       string
		expected bool
		err      bool
	}{
		{"1.0.0-alpha.1", "<", "1.0.0", true, false},
		{"1.0.0-alpha.1", ">", "1.0.0-alpha", true, false},
		{"1.0.0-alpha.1", "=", "1.0.0-alpha.1", true, false},
		{"1.0.0-beta.2", "<", "1.0.0-beta.3", true, false},
		{"1.0.0+build123", "<=", "1.0.0+build456", true, false},
		{"1.0.0", "<", "1.2.2", true, false},
		{"1.2.1", ">", "1.2.0", true, false},
		{"2.0.3", "=", "2.0.3", true, false},
		{"2.0.2", "<", "2.0.3", true, false},
		{"1.0.0-alpha.1", "<", "1.0.0", true, false},
		{"1.0.0-alpha.1", ">", "1.0.0-alpha", true, false},
		{"1.0.0-alpha.1", "=", "1.0.0-alpha.1", true, false},
		{"1.0.0-beta.2", "<", "1.0.0-beta.3", true, false},
		{"1.1.0-beta.2", ">", "1.0.0-beta.3", true, false},
		{"1.0.0-alpha", "<", "1.0.0-alpha.1", true, false},
		{"1.0.0+build123", "<=", "1.0.0+build456", true, false},
		{"1.0.0-alpha+build789", ">=", "1.0.0-alpha+build000", true, false},
		{"1.2.0-alpha", "<", "1.2.0-beta", true, false},
		{"1.2.0-alpha.1", ">", "1.2.0-alpha", true, false},
		{"1.0.0+build0001", "=", "1.0.0+build0001", true, false},
		{"1.0.0", "<", "1.2.2", true, false},
		{"1.2.1", ">", "1.2.0", true, false},
		{"2.0.3", "=", "2.0.3", true, false},
		{"2.0.2", "<", "2.0.3", true, false},
		{"1.4.3", ">", "1.0.1", true, false},
		{"1.0.0", "=", "1.0.0", true, false},
		{"1.0.0", "=", "1.0.0", true, false},
		{"1.0.0", "=", "1.0.0", true, false},
		{"1.0.0", "invalid", "1.2.4", false, true}, // Unsupported operator
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s %s", tt.v1, tt.operator, tt.v2), func(t *testing.T) {
			result, err := CompareVersions(tt.v1, tt.operator, tt.v2)
			if (err != nil) != tt.err {
				t.Errorf("expected error: %v, got: %v", tt.err, err)
			}
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestVersionEquals(t *testing.T) {
	tests := []struct {
		v1       Version
		v2       Version
		expected bool
	}{
		{Version{major: 1, minor: 2, patch: 3}, Version{major: 1, minor: 2, patch: 3}, true},
		{Version{major: 1, minor: 2, patch: 3}, Version{major: 1, minor: 2, patch: 4}, false},
		{Version{major: 1, minor: 2, patch: 3, prerelease: "alpha"}, Version{major: 1, minor: 2, patch: 3, prerelease: "alpha"}, true},
		{Version{major: 1, minor: 2, patch: 3, prerelease: "alpha"}, Version{major: 1, minor: 2, patch: 3, prerelease: "beta"}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v == %+v", tt.v1, tt.v2), func(t *testing.T) {
			result := tt.v1.Equals(tt.v2)
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestVersionLessThan(t *testing.T) {
	tests := []struct {
		v1       Version
		v2       Version
		expected bool
	}{
		{Version{major: 1, minor: 2, patch: 3}, Version{major: 1, minor: 2, patch: 4}, true},
		{Version{major: 1, minor: 2, patch: 4}, Version{major: 1, minor: 2, patch: 3}, false},
		{Version{major: 1, minor: 2, patch: 3}, Version{major: 1, minor: 3, patch: 0}, true},
		{Version{major: 1, minor: 3, patch: 0}, Version{major: 1, minor: 2, patch: 3}, false},
		{Version{major: 1, minor: 2, patch: 3, prerelease: "alpha"}, Version{major: 1, minor: 2, patch: 3}, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v < %+v", tt.v1, tt.v2), func(t *testing.T) {
			result := tt.v1.LessThan(tt.v2)
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func TestVersionGreaterThan(t *testing.T) {
	tests := []struct {
		v1       Version
		v2       Version
		expected bool
	}{
		{Version{major: 1, minor: 2, patch: 3}, Version{major: 1, minor: 2, patch: 4}, false},
		{Version{major: 1, minor: 2, patch: 4}, Version{major: 1, minor: 2, patch: 3}, true},
		{Version{major: 1, minor: 3, patch: 0}, Version{major: 1, minor: 2, patch: 3}, true},
		{Version{major: 1, minor: 2, patch: 3}, Version{major: 1, minor: 3, patch: 0}, false},
		{Version{major: 1, minor: 2, patch: 3, prerelease: "alpha"}, Version{major: 1, minor: 2, patch: 3}, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%+v > %+v", tt.v1, tt.v2), func(t *testing.T) {
			result := tt.v1.GreaterThan(tt.v2)
			if result != tt.expected {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}
