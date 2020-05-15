package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Planet struct {
	Nome    string `json:"nome"`
	Clima   string `json:"clima"`
	Terreno string `json:"terreno"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/planet", PlanetHandler).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "dbapi"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Conectou Suave")
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return db
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("suave")
}

func PlanetHandler(w http.ResponseWriter, r *http.Request) {
	b, error := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if error != nil {
		http.Error(w, error.Error(), 500)
		return
	}

	var p Planet
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("aqui 2")
		log.Fatal(err)
	}

	output, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.Write(output)

	fmt.Println(p)


	db := dbConn()
	email := p.Email
	comentario := p.Comentario
	insForm, err := db.Prepare("INSERT INTO comentario(email, comentario) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(email, comentario)
	log.Println("INSERT: Email: " + email + " | Comentario: " + comentario)

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}