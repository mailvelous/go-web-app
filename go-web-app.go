package main

import (
	"fmt"
	"net/http"
)

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome"
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world"
	w.Write([]byte(message))
}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
    http.HandleFunc("/hello", handlerHello)

    fmt.Println("Starting server on :8080...")
    err := http.ListenAndServe(":8080", nil)
    if err!= nil {
        panic(err)
    }
}