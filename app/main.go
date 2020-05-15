package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	// r.HandleFunc("/comentario", CommentHandler).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("suave")
}
