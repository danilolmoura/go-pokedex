package main

import (
	"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Pokemon struct {
    ID string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
    Type *Type `json:"type,omitempty"`
}

type Type struct {
    ID  string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}

var pokedex []Pokemon

// função principal
func main() {	
	pokedex = append(pokedex, Pokemon{ID: "1", Name: "Bulbasaur", Type: &Type {ID: "1", Name: "Water"}})
	pokedex = append(pokedex, Pokemon{ID: "2", Name: "Charmander", Type: &Type {ID: "2", Name: "Fire"}})
	pokedex = append(pokedex, Pokemon{ID: "3", Name: "Gengar"})

	router := mux.NewRouter()
	router.HandleFunc("/", Index).Methods("GET")
	router.HandleFunc("/pokedex", GetPokemons).Methods("GET")
	router.HandleFunc("/pokedex/{id}", GetPokemon).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Welcome to Go Pokedex")
}


func GetPokemons(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pokedex)
}

func GetPokemon(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range pokedex {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Pokemon{})
}