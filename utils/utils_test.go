package utils

import "testing"

func TestFormatName(t *testing.T) {
	result := FormatName("John", "Doe")
	if result != "John Doe" {
		t.Errorf("Expected 'John Doe', got %s", result)
	}
}