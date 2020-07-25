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
	r.HandleFunc("/insertplanet", route.InsertPlanet).Methods("POST")
	r.HandleFunc("/listallplanet", route.ListAllPlanet).Methods("GET")
	r.HandleFunc("/getplanetname/{nome}", route.GetNamePlanet).Methods("GET")
	r.HandleFunc("/gettplanetid/{id}/", route.GetIdPlanet).Methods("GET")
	r.HandleFunc("/deleteplanet/{id}", route.RemovePlanet).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))

}
