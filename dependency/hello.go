package main

import (
	"fmt"
	"io"
	"os"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

func Hello(n string, l string) string {
	if n == "" {
		n = "world"
	}

	return greetingPrefix(l) + n
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello %s", name)
}

func greetingPrefix(l string) (prefix string) {
	switch l {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	Greet(os.Stdout, "Elodie")
}
