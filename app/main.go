package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
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
	r.HandleFunc("/listplanet", ListAllPlanet).Methods("GET")
	r.HandleFunc("/listplanetname", ListNamePlanet).Methods("POST")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8090", r))

}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "starwars"
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
	nome := p.Nome
	clima := p.Clima
	terreno := p.Terreno
	insForm, err := db.Prepare("INSERT INTO planet(nome, clima, terreno) VALUES(?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(nome, clima, terreno)
	log.Println("INSERT: Nome: " + nome + " | Clima: " + clima + " | Terreno: " + terreno)

	defer db.Close()
	http.Redirect(w, r, "/", 301)
}



func ListAllPlanet(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
    rows, err := db.Query("SELECT nome FROM planet")
    if err != nil {
        panic(err.Error())
	}
	for rows.Next() {
        var (
            nome string
        )
        if err := rows.Scan(&nome); err != nil {
            panic(err)
        }
        fmt.Printf("%v\n", nome)
    }
    if err := rows.Err(); err != nil {
        panic(err)
	}
	
	defer dbConn()

}

func ListNamePlanet(w http.ResponseWriter, r *http.Request) {
	b, error := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if error != nil {
		http.Error(w, error.Error(), 500)
		return
	}

	var p Planet
	err := json.Unmarshal(b, &p)
	if err != nil {
		fmt.Println("aqui 1")
		log.Fatal(err)
	}


	db := dbConn()
	pName := p
    rows, err := db.Query("SELECT * FROM planet WHERE nome='?'", pName)
    if err != nil {
        panic(err.Error())
	}
	for rows.Next() {
        var (
            nome string
        )
        if err := rows.Scan(&nome); err != nil {
            panic(err)
        }
        fmt.Printf("%v\n", nome)
    }
    if err := rows.Err(); err != nil {
        panic(err)
    } 

}