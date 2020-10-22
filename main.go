package main

import (
	"encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Pokemon struct {
    ID string `json:"id"`
    Num string `json:"num"`
    Name string `json:"name"`
    Img string `json:"img"`
    Type []string `json:"type"`
    Height string `json:"heigh"`
    Weight string `json:"weigh"`
    Candy string `json:"candy"`
    CandyCount int `json:"candy_count"`
    Egg string `json:"egg"`
    SpawnChance float32 `json:"spawn_chance"`
    AvgSpawns float32 `json:"avg_spawns"`
    SpawnTime string `json:"spawn_time"`
    Multipliers []float32 `json:"multipliers"`
    Weaknesses []string `json:"weaknesses"`
    NextEvolution []*NextEvolution `json:"next_evolution"`

}

type NextEvolution struct {
    Num string `json:"num"`
    Name string `json:"name"`
}

var pokedex []Pokemon

func main() {	
	pokedex = append(
        pokedex,
        Pokemon{
            ID: "1",
            Num: "001",
            Name: "Bulbasaur",
            Img: "http://www.serebii.net/pokemongo/pokemon/001.png",
            Type: []string{
                "Grass",
                "Poison",
            },
            Height: "0.71 m",
            Weight: "6.9 kg",
            Candy: "Bulbasaur Candy",
            CandyCount: 25,
            Egg: "2 km",
            SpawnChance: 0.69,
            AvgSpawns: 69,
            SpawnTime: "20:00",
            Multipliers: []float32{
                1.58,
            },
            Weaknesses: []string{
                "Fire",
                "Ice",
                "Flying",
                "Psychic",
            },
            NextEvolution: []*NextEvolution{
                &NextEvolution{
                    Num: "002",
                    Name: "Ivysaur",
                },
                &NextEvolution{
                    Num: "003",
                    Name: "Venusaur",
                },
            },
        },
    )

    pokedex = append(
        pokedex,
        Pokemon{
            ID: "4",
            Num: "004",
            Name: "Charmander",
            Img: "http://www.serebii.net/pokemongo/pokemon/004.png",
            Type: []string{
                "Fire",
            },
            Height: "0.61 m",
            Weight: "8.5 kg",
            Candy: "Charmander Candy",
            CandyCount: 25,
            Egg: "2 km",
            SpawnChance: 0.253,
            AvgSpawns: 25.3,
            SpawnTime: "08:45",
            Multipliers: []float32{
                1.65,
            },
            NextEvolution: []*NextEvolution{
                &NextEvolution{
                    Num: "005",
                    Name: "Charmeleon",
                },
                &NextEvolution{
                    Num: "006",
                    Name: "Charizard",
                },
            },
        },
    )

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