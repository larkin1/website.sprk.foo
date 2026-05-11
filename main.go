package main

import (
	"fmt"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("html/index.html")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Fprint(w, string(file))
}

func second(w http.ResponseWriter, r *http.Request) {
	file, err := os.ReadFile("html/second.html")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Fprint(w, string(file))
}

func widget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<div class="card">Hello</div>`)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/second", second)
	mux.HandleFunc("/widget", widget)

	http.ListenAndServe(":8080", mux)
}
