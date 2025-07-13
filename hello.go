package main

import "fmt"

// Constants make easier to improve performance and to
// refactoring.
const englishGreetingWord = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return fmt.Sprint(englishGreetingWord + name)
}

func main() {
	name := "João"
	greeting := Hello(name)
	fmt.Println(greeting)
}
