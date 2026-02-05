package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	spanishHelloPrefix = "Hola, "
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("Toli", "French"))
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
