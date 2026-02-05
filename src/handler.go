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

func ConnexionHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/connexion.html")

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

	http.Redirect(w, r, "/", http.StatusFound)
}
