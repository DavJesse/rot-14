package main

import "testing"

func TestStringContains(t *testing.T) {
	testStr1, testSubStr := "Contains", "ins"
	got := stringContains(testStr1, testSubStr)
	expected := true

	if got != expected {
		t.Errorf("Got: %t", got)
		t.Errorf("Expected: %t", expected)
		t.Errorf("TestStringContains Failed!")
	}
}
