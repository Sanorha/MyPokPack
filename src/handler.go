package pokemon

import (
	"fmt"
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

func ConnexionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/connexion.html")

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

func SubmitConnexionHandler(w http.ResponseWriter, r *http.Request, jeux *Jeux) {
	pseudo := r.FormValue("pseudo")
	motdepasse := r.FormValue("mdp")
	println("Pseud et mdp : ", pseudo, motdepasse)

	motdepasse_sql := SearchSQL(pseudo)

	var check_mdp bool = CompareMDP(motdepasse, motdepasse_sql)

	fmt.Println("Bool : ", check_mdp)

	///////////////////////////////////////////////////////////
	// Gestion comparaison mdp et redirect, et message html  //
	///////////////////////////////////////////////////////////
	http.Redirect(w, r, "/", http.StatusFound)
}

func RetourHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func SubmitDeconnexionHandler(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w)

	http.Redirect(w, r, "/", http.StatusFound)
}
