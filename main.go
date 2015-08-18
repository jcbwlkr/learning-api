package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
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

	router := httprouter.New()

	router.GET("/", index)
	router.GET("/posts", list)
	router.POST("/posts", create)
	router.GET("/posts/:id", fetch)
	router.PUT("/posts/:id", update)
	router.DELETE("/posts/:id", del)

	addr := fmt.Sprintf("localhost:%d", port)
	log.Printf("Serving on %s", addr)
	log.Fatal(http.ListenAndServe(addr, router))
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

func idFromParams(ps httprouter.Params) (int, error) {
	idStr := ps.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logRequest(r)
	fmt.Fprintln(w, "Welcome to the Posts test API!")
}

func list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	serveJSON(w, db.FindAll(), http.StatusOK)
}

func create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
}

func fetch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	var id int
	var err error
	if id, err = idFromParams(ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	post, err := db.FindOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		serveJSON(w, post, http.StatusOK)
	}
}

func update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)

	var id int
	var err error
	if id, err = idFromParams(ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ = id
}

func del(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logRequest(r)
	var id int
	var err error
	if id, err = idFromParams(ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
