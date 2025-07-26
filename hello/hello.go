package main

import "fmt"

// Constants make easier to improve performance and to
// refactoring.
const (
	portuguese = "Portuguese"
	french     = "French"

	englishGreetingWord    = "Hello, "
	portugueseGreetingWord = "Olá, "
	frenchGreetingWord     = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greatingsPrefix(language) + name + "!"
}

func greatingsPrefix(language string) (prefix string) {
	switch language {
	case portuguese:
		prefix = portugueseGreetingWord
	case french:
		prefix = frenchGreetingWord
	default:
		prefix = englishGreetingWord
	}
	return prefix
}

func main() {
	name := "João"
	greeting := Hello(name, "")
	fmt.Println(greeting)
}
