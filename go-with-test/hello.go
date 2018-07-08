// Hello file

package main

import (
	"fmt"
	)

// Refractor by using constant

const helloPrefix = "Hello, "

func Hello(name string) string{
	return helloPrefix + name
}

func main(){
	fmt.Println(Hello("Chris"))
}


