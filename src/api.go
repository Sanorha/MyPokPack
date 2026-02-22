package pokemon

import (
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func RandBooster() (string, string, string) {
	api := "https://pokeapi.co/api/v2/pokemon/"

	url := api + strconv.Itoa(rand.IntN(1025))

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	type Data struct {
		Name string `json:"name"`

		Types []struct {
			Type struct {
				Name string `json:"name"`
			} `json:"type"`
		} `json:"types"`

		Sprites struct {
			Other struct {
				Dream_world struct {
					Image string `json:"front_default"`
				} `json:"dream_world"`
			} `json:"other"`
		} `json:"sprites"`
	}

	data := Data{}

	json.NewDecoder(resp.Body).Decode(&data)

	var types string

	for _, t := range data.Types {
		types += (t.Type.Name + " ")
	}

	pokemon_name := data.Name
	pokemon_types := types
	pokemon_image := data.Sprites.Other.Dream_world.Image

	return pokemon_name, pokemon_types, pokemon_image
}
