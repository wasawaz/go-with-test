package main

import (
	"fmt"
)

var langeHelloPrefix = map[string]string{
	"English": "Hello",
	"Spanish": "Hola",
}

const english = "english"
const spanish = "spanish"
const french = "french"

const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := getGreetingPrefix(language)
	return fmt.Sprintf("%s, %s", prefix, name)
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}

func main() {
	fmt.Println(Hello("world", "english"))
}
