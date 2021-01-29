package main

import (
	"net/http"
	"io"
	"fmt"
)

//Greet used to write greeting words to the buffer
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

//MyGreeterHandler send message to the response stream
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world!")
}

func main() {
	//Greet(os.Stdout, "Lyle")
	http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler))
}