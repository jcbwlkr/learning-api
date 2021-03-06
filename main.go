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

	db.Insert(Article{User: "jane", Body: "Hello, world!"})
	db.Insert(Article{User: "john", Body: "Lorem ipsum dolor sit amet."})
}

func main() {
	flag.Parse()

	router := httprouter.New()

	router.GET("/", index)
	router.GET("/articles", list)
	router.POST("/articles", create)
	router.OPTIONS("/articles", options)
	router.GET("/articles/:id", fetch)
	router.PUT("/articles/:id", update)
	router.DELETE("/articles/:id", del)
	router.OPTIONS("/articles/:id", options)
	router.ServeFiles("/site/*filepath", http.Dir(""))

	addr := fmt.Sprintf("localhost:%d", port)

	fmt.Printf("\nBlog articles test API!\n\n")

	fmt.Printf("This server is now listening on %s\n", addr)
	fmt.Printf("If you ran this command from inside your site's folder you can view your site at http://%s/site/\n", addr)

	fmt.Println("You can make the following API requests")
	fmt.Printf("GET    http://%s/articles      -- Get all articles\n", addr)
	fmt.Printf("POST   http://%s/articles      -- Make an article. Send a body like {\"user\": \"alice\", \"body\": \"foo\"}\n", addr)
	fmt.Printf("GET    http://%s/articles/:id  -- Get a particular article\n", addr)
	fmt.Printf("PUT    http://%s/articles/:id  -- Update an article. Send a body like {\"user\": \"anna\", \"body\": \"bar\"}\n", addr)
	fmt.Printf("DELETE http://%s/articles/:id  -- Delete an article\n", addr)

	fmt.Printf("\nPress Ctrl-c at any time to kill this server\n\n")

	var h http.Handler = router
	h = mwCORSHeaders(h)
	h = mwLogger(h)

	log.Fatal(http.ListenAndServe(addr, h))
}

func mwLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %-6s %s\n", r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func mwCORSHeaders(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	})
}

func serveJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(&data); err != nil {
		s := http.StatusInternalServerError
		http.Error(w, http.StatusText(s), s)
		return
	}
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
	fmt.Fprintln(w, "Welcome to the Articles test API!")
}

func options(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Methods", "PUT, POST, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.WriteHeader(http.StatusOK)
}

func list(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	serveJSON(w, db.FindAll(), http.StatusOK)
}

func create(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	article := Article{}
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article = db.Insert(article)

	serveJSON(w, article, http.StatusCreated)
}

func fetch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		id  int
		err error
	)
	if id, err = idFromParams(ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if article, err := db.FindOne(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	} else {
		serveJSON(w, article, http.StatusOK)
	}
}

func update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		id  int
		err error
	)
	if id, err = idFromParams(ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article, err := db.FindOne(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Update(article)
	serveJSON(w, article, http.StatusOK)
}

func del(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var (
		id  int
		err error
	)
	if id, err = idFromParams(ps); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Delete(id)
	w.WriteHeader(http.StatusNoContent)
}
