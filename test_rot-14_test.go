package main

import "testing"

func TestToRot14(t *testing.T) {
	testStr := "Hello! How are You?"
	got := toRot14(testStr)
	expected := "Vszzc! Vck ofs Mci?"

	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestRot14 failed")
	}
}

func TestFromRot14(t *testing.T) {
	testStr := "Vszzc! Vck ofs Mci?"
	got := toRot14(testStr)
	expected := "Hello! How are You?"

	if got != expected {
		t.Errorf("Got: %s", got)
		t.Errorf("Expected: %s", expected)
		t.Errorf("TestRot14 failed")
	}
}