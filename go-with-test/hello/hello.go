// Hello file

package main

import (
	"fmt"
)

// Refractor by using constant

const helloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return helloPrefix + name
}

func main() {
	fmt.Println(Hello("Chris"))
}
