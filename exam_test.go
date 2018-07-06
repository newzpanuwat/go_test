package main
import "testing"


result := sum(2, 3)

func TestGo(t *testing.T) {
	if sum != 5 {
		t.Errorf("error expect 5, but got '%d' ", result)		
	}
}

func sum(arg1, arg2 int) int{
	return arg1 + arg2
}