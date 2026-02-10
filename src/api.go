package pokemon

import (
	"encoding/json"
	"log"
	"net/http"
)

func Affiche() {
	resp, err := http.Get("https://pokeapi.co/docs/v2#pokemon")

	if err != nil {
		log.Fatal(err)
	}
	var jeux []Jeux

	json.NewDecoder(resp.Body).Decode(&jeux)
	println(resp.Body)
}
