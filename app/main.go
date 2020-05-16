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

type Api struct {
	Name  string `json:"name"`
	Films string `json:"films"`
}

type Planet struct {
	Id      uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Nome    string `json:"nome"`
	Clima   string `json:"clima"`
	Terreno string `json:"terreno"`
}

// main Star Wars Planet Search
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/planet", PlanetHandler).Methods("POST")
	r.HandleFunc("/listplanet", ListAllPlanet).Methods("GET")
	r.HandleFunc("/listplanetname/{nome}", ListNamePlanet).Methods("GET")
	r.HandleFunc("/listplanetid/{id}", ListIdPlanet).Methods("GET")
	r.HandleFunc("/deleteplanet/{id}", RemovePlanet).Methods("DELETE")
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
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	db := dbConn()
	rows, err := db.Query("SELECT nome FROM planet")
	if err != nil {
		panic(err.Error())
	}

	var planet Planet

	for rows.Next() {
		var (
			nome string
		)
		if err := rows.Scan(&nome); err != nil {
			panic(err)
		}
		fmt.Printf("%v\n", nome)
	}
	json.NewEncoder(w).Encode(planet)

	if err := rows.Err(); err != nil {
		panic(err)
	}

	defer dbConn()

}

func ListNamePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	var nome = params["nome"]
	db := dbConn()
	result, err := db.Query("SELECT * FROM planet WHERE nome=?", nome)
	if err != nil {
		log.Panic(err)
	}

	var planet Planet
	for result.Next() {
		err := result.Scan(&planet.Id, &planet.Nome, &planet.Clima, &planet.Terreno)
		if err != nil {
			panic(err.Error())
		}
	}

	json.NewEncoder(w).Encode(planet)

}

func ListIdPlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	var id = params["id"]
	db := dbConn()
	result, err := db.Query("SELECT * FROM planet WHERE id=?", id)
	if err != nil {
		log.Panic(err)
	}

	var planet Planet
	for result.Next() {
		err := result.Scan(&planet.Id, &planet.Nome, &planet.Clima, &planet.Terreno)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(planet)
}

func RemovePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	db := dbConn()
	stmt, err := db.Prepare("DELETE FROM planet WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Planet with ID = %s was deleted", params["id"])
}

func GetApiSW(w http.ResponseWriter, r *http.Request) {
	// get api
	resp, err := http.Get("https://swapi.dev/api/")
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	var api Api
	body, err := ioutil.ReadAll(resp.Body)
	error := json.Unmarshal(body, &api)
	if error != nil {
		fmt.Println("aqui 3")
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(api)
}
