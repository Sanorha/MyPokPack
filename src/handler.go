package pokemon

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func InscriptionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/inscription.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func ConnectionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/connection.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func SubmitInscriptionHandler(w http.ResponseWriter, r *http.Request) {
	pseudo := r.FormValue("pseudo")
	motdepasse := r.FormValue("mdp")
	email := r.FormValue("email")

	motdepasse_hash, err := CreateHashMDP(motdepasse)

	if err != nil {
		log.Fatal(err)
	}

	motdepasse = motdepasse_hash

	println(pseudo, email, motdepasse)

	AddUser(pseudo, email, motdepasse)
	AddCookie(w, pseudo)

	http.Redirect(w, r, "/", http.StatusFound)
}

func SubmitConnectionHandler(w http.ResponseWriter, r *http.Request) {
	pseudo := r.FormValue("pseudo")
	motdepassehash := r.FormValue("mdp")
	println(pseudo, motdepassehash)
	http.Redirect(w, r, "/", http.StatusFound)

	///////////////////////////////////////////////////////////
	// REQUETE SQL POUR recup INFO base donnée pseudo		 //
	///////////////////////////////////////////////////////////

	// si pseudo existe

	///////////////////////////////////////////////////////////
	// REQUETE SQL POUR recup INFO base donnée mdp //
	///////////////////////////////////////////////////////////
}

func RetourHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func SubmitDeconnectionHandler(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w)

	http.Redirect(w, r, "/", http.StatusFound)
}
