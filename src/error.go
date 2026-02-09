package pokemon

import (
	"log"
)

func ReadError(err error) {
	if err != nil {
		log.Println("Erreur lors de l'inscription& :", err)

		return
	}
}
