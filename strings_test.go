package main

import "testing"

func TestStringContains(t *testing.T) {
	// Define test input values
	testStr1, testSubStr := "Contains", "ins"

	// Call stringContains for testing
	got := stringContains(testStr1, testSubStr)

	// Define the expected output
	expected := true

	// Check if the output matches the expected result
	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestStringContains Failed!")
	}
}
