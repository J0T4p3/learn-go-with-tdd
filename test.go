package main

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %v!", name)
}

func main() {
	name := "João"

	greeting := Hello(name)

	fmt.Println(greeting)
}
