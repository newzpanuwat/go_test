package Iteration

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, repeated, expected string) {
		t.Helper()
		if repeated != expected {
			fmt.Errorf("expected '%s' but got '%s'", repeated, expected)
		}
	}

	t.Run("Test first iteration", func(t *testing.T) {
		repeated := Repeat5("a")
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})

	t.Run("Test second iteration", func(t *testing.T) {
		repeated := Repeat10("x")
		expected := "xxxxxxxxxx"
		assertCorrectMessage(t, repeated, expected)
	})
}

// Benchmark func for testing
func BenchmarkRepeat5times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat5("a")
	}
}
func BenchmarkRepeat10times(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat10("x")
	}
}
