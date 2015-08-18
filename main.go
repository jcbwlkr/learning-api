package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	db   DB
	port int
)

func init() {
	flag.IntVar(&port, "port", 8080, "The port to bind to on localhost")

	db.Insert(Post{User: "jane", Message: "Hello, world!"})
	db.Insert(Post{User: "john", Message: "Lorem ipsum dolor sit amet."})
}

func main() {
	flag.Parse()

	log.Printf("Serving on localhost:%d", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		fmt.Fprintln(w, "Welcome to the Posts test API!")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func logRequest(r *http.Request) {
	log.Printf("Serving %s\n", r.URL.Path)
}
