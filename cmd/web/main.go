package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "HTTP network address")
	flag.Parse()
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(http.Dir("M:/Projects/Snippet_box/ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	
	log.Printf("Starting server on: %v\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}