package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "

	spanish = "Spanish"
	french  = "French"

	defaultNameSuffix = "World"
)

func greetingPrefix(lang string) string {
	var prefix string

	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func Hello(name string, lang string) string {
	if name == "" {
		name = defaultNameSuffix
	}

	return greetingPrefix(lang) + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
