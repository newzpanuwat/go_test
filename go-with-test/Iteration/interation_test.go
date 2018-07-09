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
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
}

// Benchmark func for testing
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
