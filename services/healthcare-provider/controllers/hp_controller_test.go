package controllers

import "testing"

func TestComparePasswords(t *testing.T) {

	// Add user argument
	got, _ := ComparePasswords("10", "10")
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}
