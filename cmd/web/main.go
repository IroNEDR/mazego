package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func main() {

	r := chi.NewRouter()
	r.Get("/", getIndex)
}
