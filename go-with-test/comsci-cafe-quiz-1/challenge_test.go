package main

import "testing"

func TestChallenge(t *testing.T) {
	got := Challenge(1)
	want := 450

	if got != want {
		t.Errorf("GOT '%d', WANT '%d'", got, want)
	}
}
