package main

import "testing"

func TestQuiz(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s', Want '%s'", got, want)
		}
	}

	t.Run("testing hello for string", func(t *testing.T) {
		got := Hello("Newhales")
		want := "Hello, Newhales"
		assertCorrectMessage(t, got, want)

	})

	t.Run("testing hello for empty string", func(t *testing.T) {
		got := Hello("")
		want := "Hello, Empty"
		assertCorrectMessage(t, got, want)

	})
}

func TestAddition(t *testing.T) {

	insertCorrectMessage := func(t *testing.T, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("Testing add function 2 + 2", func(t *testing.T) {
		got := add(2, 3)
		want := 5
		insertCorrectMessage(t, got, want)
	})
}
