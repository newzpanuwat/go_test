package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}


func main () {
	fmt.Println("Hello, Gophers!!!")
	fmt.Println(add(40, 10))
}
