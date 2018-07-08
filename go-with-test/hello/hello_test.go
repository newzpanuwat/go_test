//Testing file

package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("Got '%s' want '%s'", got, want)
		}
	}

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in English", func(t *testing.T) {
		got := Hello("David Beackham", "English")
		want := "Hello, David Beackham"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Mbappe", "French")
		want := "Bonjour, Mbappe"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in default with empty string", func(t *testing.T) {
		got := Hello("Frank Lampard", "")
		want := "Hello, Frank Lampard"
		assertCorrectMessage(t, got, want)
	})

}
