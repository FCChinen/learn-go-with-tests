package main

import "fmt"

const french = "French"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	m := map[string]string {
		"French": "Bonjour, ",
		"Spanish": "Hola, ",
	}
	prefix = m[language]
	if prefix == "" {
		return "Hello, "
	}
	return 
}

func main() {
	fmt.Println(Hello("world", ""))
}
