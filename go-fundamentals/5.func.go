package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func divide(x int, y int) int {
	return x / y
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(add(1, 5))
}

