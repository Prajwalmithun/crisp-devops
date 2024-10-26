package main

import (
    "fmt"
    "net/http"
)

func main() {
	fmt.Println("HTTP Server Running.......")
    http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "pong from Go service")
    })
    http.ListenAndServe(":8888", nil)
}