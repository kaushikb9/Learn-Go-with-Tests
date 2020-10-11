package main

import "fmt"

var prefix string

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case "Spanish":
		prefix = spanishHelloPrefix
	case "French":
		prefix = frenchHelloPrefix
	case "English":
		prefix = englishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return getGreetingPrefix(language) + name
}

func main() {
	name := "Chris"
	fmt.Println(Hello(name, ""))
}
