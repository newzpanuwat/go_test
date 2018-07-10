package challenges

import "testing"

func TestConvert(t *testing.T) {
	got := Convert(1)
	want := 37.5

	if got != want {
		t.Errorf("got '%f', want '%f'", got, want)
	}
}
