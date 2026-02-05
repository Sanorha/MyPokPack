package pokemon

import (
	"fmt"
	"net/http"
)

func Server() {
	user := InitUser()

	// affiche index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r)
	})

	// affiche inscription.html
	http.HandleFunc("/inscription", func(w http.ResponseWriter, r *http.Request) {
		InscriptionHandler(w, r)
	})

	// affiche connection.html
	http.HandleFunc("/connection", func(w http.ResponseWriter, r *http.Request) {
		ConnectionHandler(w, r)
	})

	// recup, pseudo, mail, mdp, hash, créer cookie, redirige sur /
	http.HandleFunc("/submit_inscription", func(w http.ResponseWriter, r *http.Request) {
		SubmitInscriptionHandler(w, r, &user)
	})

	// recup, pseudo, mdp, compare hash et mdp, créer cookie, redirige sur
	http.HandleFunc("/submit_connexion", func(w http.ResponseWriter, r *http.Request) {
		SubmitConnectionHandler(w, r)
	})

	// supprime le cookie, redirige sur /
	http.HandleFunc("/deconnexion", func(w http.ResponseWriter, r *http.Request) {
		SubmitDeconnectionHandler(w, r)
	})

	// bontou retour redirige sur /
	http.HandleFunc("/retour", func(w http.ResponseWriter, r *http.Request) {
		RetourHandler(w, r)
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serveur lancer sur 8080")
	http.ListenAndServe(":8080", nil)
}
