package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "Empty"
	}
	return helloPrefix + name
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(Hello("Newhales"))
	fmt.Println(add(2, 4))
}
