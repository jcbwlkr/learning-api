package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
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

	http.HandleFunc("/posts", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		idStr := r.URL.Path[len("/posts/"):]

		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			switch r.Method {
			case "GET":
				post, err := db.FindOne(id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusNotFound)
				} else {
					serveJSON(w, post, http.StatusOK)
				}
			case "PUT":
			case "DELETE":
				db.Delete(id)
				w.WriteHeader(http.StatusNoContent)
			}
			return
		}

		switch r.Method {
		case "GET":
			serveJSON(w, db.FindAll(), http.StatusOK)
		case "POST":
		}

	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logRequest(r)
		fmt.Fprintln(w, "Welcome to the Posts test API!")
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", port), nil))
}

func logRequest(r *http.Request) {
	log.Printf("Serving %s\n", r.URL.Path)
}

func serveJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		s := http.StatusInternalServerError
		http.Error(w, http.StatusText(s), s)
		return
	}
	w.WriteHeader(status)
}
