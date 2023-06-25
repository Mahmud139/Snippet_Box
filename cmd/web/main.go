package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)
mux.
	fileServer := http.FileServer(http.Dir("M:/Projects/Snippet_box/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	log.Println("Starting server on: localhost:8080")
	err := http.ListenAndServe("localhost:8080", mux)
	log.Fatal(err)
}