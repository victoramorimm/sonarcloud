package main

import "testing"

func TestSum(t *testing.T) {
	result := sum(2, 3)

	if result != 4 {
		t.Error("Expected 5, got ", result)
	}
}