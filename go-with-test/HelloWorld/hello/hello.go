// Hello file

package main

import (
	"fmt"
)

// Refractor by using constant
const spanish = "Spanish"
const french = "French"
const englishPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {

	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("Elodie", "Spanish"))
	fmt.Println(Hello("Mbappe", "French"))
	fmt.Println(Hello("Beckham", ""))
	fmt.Println(Hello("Lampard", ""))
}
