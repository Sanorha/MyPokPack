package pokemon

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, check *Check, booster *Booster) {
	tmpl, err := template.ParseFiles("index.html", "pages/templates/showbooster.html")

	if err != nil {
		log.Fatal(err)
	}

	type Data struct {
		Check   Check
		Booster Booster
	}

	data := Data{
		Check:   *check,
		Booster: *booster,
	}

	tmpl.Execute(w, data)
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

func SubmitConnexionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
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

func SubmitDeconnexionHandler(w http.ResponseWriter, r *http.Request, check *Check, booster *Booster) {
	DeleteCookie(w)
	check.Check_connexion = false
	booster.Booster = []Pokemon{}
	http.Redirect(w, r, "/", http.StatusFound)
}

func OpenBoosterHandler(w http.ResponseWriter, r *http.Request, booster *Booster, check *Check) {
	booster.Booster = []Pokemon{}
	var data Pokemon

	check.Check_openbooster = true

	for i := 0; i < 5; i++ {
		pokemon_name, pokemon_types, pokemon_image := RandBooster()

		//quelque pokemon n'ont pas d'image sur l'API
		if pokemon_image == "" {
			i--
		} else {
			data = Pokemon{pokemon_name, pokemon_types, pokemon_image}
			booster.Booster = append(booster.Booster, data)

			pseudo_cookie := ReadCookie(w, r)
			AddJeuxSQL(pokemon_name, pokemon_types, pokemon_image, pseudo_cookie)
		}
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func RedirectBoosterHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	check.Check_openbooster = false
	http.Redirect(w, r, "/", http.StatusFound)
}

func ShowCollectionHandler(w http.ResponseWriter, r *http.Request, jeux *Jeux_slice) {
	pseudo_cookie := ReadCookie(w, r)

	SearchJeuxSQL(jeux, pseudo_cookie)

	http.Redirect(w, r, "/collection", http.StatusFound)
}

func CollectionHandler(w http.ResponseWriter, r *http.Request, jeux *Jeux_slice) {
	tmpl, err := template.ParseFiles("pages/collection.html", "pages/templates/showcollection.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, jeux)
}
