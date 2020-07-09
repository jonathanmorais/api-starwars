package main

import (
	"log"
	"net/http"
	"github.com/jonathanmorais/api-starwars/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// main Star Wars Planet Search
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", route.HomeHandler).Methods("GET")
	r.HandleFunc("/planet", route.PlanetHandler).Methods("POST")
	r.HandleFunc("/listplanet", route.ListAllPlanet).Methods("GET")
	r.HandleFunc("/listplanetname/{nome}", route.ListNamePlanet).Methods("GET")
	r.HandleFunc("/listplanetid/{id}/", route.ListIdPlanet).Methods("GET")
	r.HandleFunc("/deleteplanet/{id}", route.RemovePlanet).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))

}
