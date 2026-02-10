package pokemon

import (
	"fmt"
	"net/http"
)

func Server() {
	jeux := InitJeux()
	check := InitCheck()

	CreateTableone()

	// Affiche(&jeux)

	// affiche index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HomeHandler(w, r, &check)
	})

	// affiche inscription.html
	http.HandleFunc("/inscription", func(w http.ResponseWriter, r *http.Request) {
		InscriptionHandler(w, r)
	})

	// affiche connection.html
	http.HandleFunc("/connexion", func(w http.ResponseWriter, r *http.Request) {
		ConnexionHandler(w, r, &check)
	})

	// recup, pseudo, mail, mdp, hash, créer cookie, redirige sur /
	http.HandleFunc("/submit_inscription", func(w http.ResponseWriter, r *http.Request) {
		SubmitInscriptionHandler(w, r)
	})

	// recup, pseudo, mdp, compare hash et mdp, créer cookie, redirige sur
	http.HandleFunc("/submit_connexion", func(w http.ResponseWriter, r *http.Request) {
		SubmitConnexionHandler(w, r, &jeux, &check)
	})

	// supprime le cookie, redirige sur /
	http.HandleFunc("/deconnexion", func(w http.ResponseWriter, r *http.Request) {
		SubmitDeconnexionHandler(w, r, &check)
	})

	// bontou retour redirige sur /
	http.HandleFunc("/retour", func(w http.ResponseWriter, r *http.Request) {
		RetourHandler(w, r)
	})

	// bonton voir collection /
	http.HandleFunc("/collection", func(w http.ResponseWriter, r *http.Request) {
		CollectionHandler(w, r)
	})

	/////////////// handle ouvrir booster ///////////////////

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Serveur lancer sur 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
