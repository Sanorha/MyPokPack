package pokemon

import (
	"fmt"
	"net/http"
)

func Server() {
	user := InitUser()
	Jeux := InitJeux()
	CreateTableone()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r)
	})

	http.HandleFunc("/inscription", func(w http.ResponseWriter, r *http.Request) {
		InscriptionHandler(w, r)
	})

	http.HandleFunc("/submit_inscription", func(w http.ResponseWriter, r *http.Request) {
		SubmitInscriptionHandler(w, r, &user)
	})

	http.HandleFunc("/submit_connexion", func(w http.ResponseWriter, r *http.Request) {
		SubmitConnectionHandler(w, r, &Jeux)
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serveur lancer sur 127.0.0.1:8080")
	http.ListenAndServe("0.0.0.0:8080", nil)
}
