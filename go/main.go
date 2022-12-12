package main

import (
	"fmt"
	"net/http"
)

func root (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func health (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Healthy!</h1>")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/health_check", health)
	fmt.Println("Server start in progress ...")
	http.ListenAndServe(":3000", nil)
}