package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// main Star Wars Planet Search
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/planet", PlanetHandler).Methods("POST")
	r.HandleFunc("/listplanet", ListAllPlanet).Methods("GET")
	r.HandleFunc("/listplanetname/{nome}", ListNamePlanet).Methods("GET")
	r.HandleFunc("/listplanetid/{id}/", ListIdPlanet).Methods("GET")
	r.HandleFunc("/deleteplanet/{id}", RemovePlanet).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))

}
