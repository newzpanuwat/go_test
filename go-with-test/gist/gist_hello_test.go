package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello()
 	 want := "Hello, World"
  
  if got != want {
    t.Errorf("this func got '%s', want '%s'", got, want)
  }
}

