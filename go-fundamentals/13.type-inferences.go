/*

Type inferences

when declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
the variable's type is typed, is inferred from the value on right hand side.package fundamentals

i := 42           // int
f := 3.142        // float64
g := 0.867 + 0.5i // complex128

*/
package main

import "fmt"

func main() {
	s := "Hi there!!" //change me!!!
	b := true         //change me!!!
	f := 10.4         //change me!!!
	i := 56           //change me!!!
	fmt.Printf("s is of type %T\n", s)
	fmt.Printf("b is of type %T\n", b)
	fmt.Printf("f is of type %T\n", f)
	fmt.Printf("i is of type %T\n", i)
}
