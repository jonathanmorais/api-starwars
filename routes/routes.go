package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/jonathanmorais/api-starwars/database"


	"github.com/gorilla/mux"
)

type Api struct {
	Films []string `json:"films"`
}

type Film struct {
	Title string `json:"title"`
}

type Planet struct {
	Id      uint   `json:"id" sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Nome    string `json:"nome"`
	Clima   string `json:"clima"`
	Terreno string `json:"terreno"`
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

	db := database.DbConn()
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
	db := database.DbConn()
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

	defer database.DbConn()

}

func ListNamePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	var nome = params["nome"]
	db := database.DbConn()
	result, err := db.Query("SELECT * FROM planet WHERE nome=?", nome)
	if err != nil {
		log.Println(err)
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
	url := "https://swapi.dev/api/planets/"

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	var id = params["id"]
	db := database.DbConn()
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

	resp, err := http.Get(url + id + "/")

	var api Api

	b, err := ioutil.ReadAll(resp.Body)
	error := json.Unmarshal(b, &api)
	if error != nil {
		log.Panic(error)
	}

	json.NewEncoder(w).Encode(planet)
	json.NewEncoder(w).Encode(api)

}

func RemovePlanet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)
	db := database.DbConn()
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
