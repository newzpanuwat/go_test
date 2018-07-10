package main

import "testing"

func TestMultiple(t *testing.T) {
	got := Multiply(5)
	want := 10

	if got != want {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}
