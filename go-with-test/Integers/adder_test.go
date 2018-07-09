package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, sum, expected int) {
		t.Helper()
		if sum != expected {
			t.Errorf("expected '%d', but got '%d'", expected, sum)
		}
	}

	t.Run("normal adding", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5
		assertCorrectMessage(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
