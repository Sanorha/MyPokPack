package pokemon

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func Affiche(jeux *Jeux) {
	api := "https://pokeapi.co/api/v2/pokemon/"

	url := api + strconv.Itoa(rand.IntN(1025))

	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	json.NewDecoder(resp.Body).Decode(jeux)

	fmt.Print("Nom : ", jeux.Name, "\nType : ")
	for _, t := range jeux.Types {
		fmt.Print(t.Type.Name, " ")
	}
	fmt.Print("\nAdresse image : ", jeux.Sprites.Other.Dream_world.Image, " ")
}
