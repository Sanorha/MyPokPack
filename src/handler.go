package pokemon

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	tmpl, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, check)
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

func SubmitInscriptionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
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
	check.Check_connexion = true

	http.Redirect(w, r, "/", http.StatusFound)
}

func SubmitConnexionHandler(w http.ResponseWriter, r *http.Request, jeux *Jeux, check *Check) {
	check.Check_pseudo = false
	check.Check_mdp = false

	pseudo := r.FormValue("pseudo")
	motdepasse := r.FormValue("mdp")

	var motdepasse_sql string = SearchUserSQL(pseudo)

	if motdepasse_sql == "" {
		check.Check_pseudo = true
		fmt.Println("pb pseudo")
		http.Redirect(w, r, "/connexion", http.StatusFound)

	} else {
		var check_mdp bool = CompareMDP(motdepasse, motdepasse_sql)

		if check_mdp {
			AddCookie(w, pseudo)
			check.Check_connexion = true
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

func SubmitDeconnexionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	DeleteCookie(w)
	check.Check_connexion = false
	http.Redirect(w, r, "/", http.StatusFound)
}

func CollectionHandler(w http.ResponseWriter, r *http.Request) {

	/////////////// afficher collection ///////////////////
}

func BoosterHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("pages/booster.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, nil)
}

func OpenBoosterHandler(w http.ResponseWriter, r *http.Request, jeux *Jeux) {
	for i := 0; i < 7; i++ {
		Affiche(jeux)

		fmt.Print("Nom : ", jeux.Name, "\nType : ")

		fmt.Print("\nAdresse image : ", jeux.Image, " ")

		if jeux.Image == "" {
			i--
		} else {
			pseudo_cookie := ReadCookie(w, r)
			AddJeuxSQL(jeux, pseudo_cookie)
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
	//// fonction pour carte apparaisse sur booster.html////
}

/////////////// handler ouvrir ///////////////////
/////////////// ajout base sql ///////////////////
