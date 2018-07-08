//Testing file

package main

import "testing"

func TestHello(t *testing.T){

	t.Run("saying hello to people", func(t. *testing.T){
	 got := Hello("Chris")
         want := "Hello Chris"

         if got != want {t.Errorf("got '%s' want '%s'", got, want)}
	 
	})

	
        t.Run("saying hello world when an empty string is supplied", func(t. *testing.T){
         got := Hello("")
	 want := "Hello, World"

         if got != want {t.Errorf("got '%s' want '%s'", got, want)}

        })

}
