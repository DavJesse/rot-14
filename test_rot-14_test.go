package main

import "testing"

func TestRot14(t *testing.T) {
	testStr := "Hello! How are You?"
	got := toRot14(testStr)
	expected := "Vszzc! Vck ofs Mci?"

	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestRot14 failed")
	}
}
