package main

import "testing"

func TestRot14(t *testing.T) {
	testStr := "abc"
	got := rot14(testStr)
	expected := "opq"

	if got != expected {
		t.Errorf("Got: ", got)
		t.Errorf("Expected: ", expected)
		t.Errorf("TestRot14 failed")
	}

}