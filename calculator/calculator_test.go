package calculator

import "testing"

func TestAdd(t *testing.T) {
	result := Add(10, 5)
	if result != 15 {
		t.Errorf("Expected 15, got %d", result)
	}
}

func TestDivide(t *testing.T) {
	result := Divide(10, 2)
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}