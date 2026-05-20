package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/ascii", func(response http.ResponseWriter, m *http.Request) {
		fmt.Fprintln(response, "Welcome to ASCII Art Web")
	})
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/html", html)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./frontend/a.html")
	})

	port := ":8080"
	fmt.Println("server is running on port:8080")

	log.Fatal(http.ListenAndServe(port, nil))

}

func hello(r http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(r, "Hello, Celebrate")
}

func html(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte("<h1>Hi, I'm HTML in Go.</h1>"))
}
