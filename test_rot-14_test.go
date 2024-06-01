package main

import "testing"

func TestToRot14(t *testing.T) {
	// Define test input value
	testStr := "Hello! How are You?"

	// Call toRot14 function for testing
	got := toRot14(testStr)

	// Define the expected output
	expected := "Vszzc! Vck ofs Mci?"

	// Check if the output matches the expected result
	if got != expected {

		// If not, print error messages
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestRot14 failed")
	}
}

func TestFromRot14(t *testing.T) {
	// Define test input value
	testStr := "Vszzc! Vck ofs Mci?"

	// Call fromRot14 function for testing
	got := fromRot14(testStr)

	// Define the expected output
	expected := "Hello! How are You?"

	// Check if the output matches the expected result
	if got != expected {

		// If not, print error messages
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestRot14 failed")
	}
}
