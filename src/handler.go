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

func ConnexionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	tmpl, err := template.ParseFiles("pages/connexion.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, check)
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

func SubmitConnexionHandler(w http.ResponseWriter, r *http.Request, jeux *Jeux, check *Check) {
	check.Check_pseudo = false
	check.Check_mdp = false

	pseudo := r.FormValue("pseudo")
	motdepasse := r.FormValue("mdp")

	var motdepasse_sql string = SearchSQL(pseudo)

	if motdepasse_sql == "" {
		check.Check_pseudo = true
		fmt.Println("pb pseudo")
		http.Redirect(w, r, "/connexion", http.StatusFound)

	} else {
		var check_mdp bool = CompareMDP(motdepasse, motdepasse_sql)

		if check_mdp {
			AddCookie(w, pseudo)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			check.Check_mdp = true
			http.Redirect(w, r, "/connexion", http.StatusFound)
		}
	}
}

func RetourHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func SubmitDeconnexionHandler(w http.ResponseWriter, r *http.Request) {
	DeleteCookie(w)

	http.Redirect(w, r, "/", http.StatusFound)
}
