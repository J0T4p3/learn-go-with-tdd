package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// With the implementation of io.Writer, we can send the
// content to whenever output we can think of.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// For example, for web
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	Greet(os.Stdout, "João")
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
