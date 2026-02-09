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

func SubmitInscriptionHandler(w http.ResponseWriter, r *http.Request, user *User) {
	user.Pseudo = r.FormValue("pseudo")
	user.Motdepasse = r.FormValue("mdp")
	user.Email = r.FormValue("email")
	println(user.Pseudo, user.Motdepasse, user.Email)
	TableUser(&user.Pseudo, &user.Motdepasse, &user.Email)

	AddCookie(w)

	///////////////////////////////////////////////////////
	// REQUETE SQL POUR envoyer INFO sur base donnée	 //
	///////////////////////////////////////////////////////

	http.Redirect(w, r, "/", http.StatusFound)
}

func SubmitConnectionHandler(w http.ResponseWriter, r *http.Request, jeux *Jeux) {
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
