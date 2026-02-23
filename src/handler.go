package pokemon

import (
	"log"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request, check *Check, booster *Booster) {
	check.Check_pseudo_inscription = false
	check.Check_email_inscription = false
	check.Check_pseudo_connection = false
	check.Check_mdp_connection = false

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

func InscriptionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	tmpl, err := template.ParseFiles("pages/inscription.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, check)
}

func ConnexionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	tmpl, err := template.ParseFiles("pages/connexion.html")

	if err != nil {
		log.Fatal(err)
	}

	tmpl.Execute(w, check)
}

func SubmitInscriptionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	check.Check_pseudo_inscription = false
	check.Check_email_inscription = false

	pseudo := r.FormValue("pseudo")
	motdepasse := r.FormValue("mdp")
	email := r.FormValue("email")

	var pseudo_sql string = SearchUserSQL(pseudo)
	var email_sql string = SearchEmailSQL(email)

	// pseudo ou mail deja existant
	if pseudo_sql == pseudo || email_sql == email {
		if pseudo_sql == pseudo {
			check.Check_pseudo_inscription = true
		}
		if email_sql == email {
			check.Check_email_inscription = true
		}
		http.Redirect(w, r, "/inscription", http.StatusFound)

	} else {
		motdepasse_hash, err := CreateHashMDP(motdepasse)

		if err != nil {
			log.Fatal(err)
		}

		motdepasse = motdepasse_hash

		AddUser(pseudo, email, motdepasse)
		AddCookie(w, pseudo)
		check.Check_connexion = true

		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func SubmitConnexionHandler(w http.ResponseWriter, r *http.Request, check *Check) {
	check.Check_pseudo_connection = false
	check.Check_mdp_connection = false

	pseudo := r.FormValue("pseudo")
	motdepasse := r.FormValue("mdp")

	var motdepasse_sql string = SearchMdpSQL(pseudo)

	if motdepasse_sql == "" {
		check.Check_pseudo_connection = true
		http.Redirect(w, r, "/connexion", http.StatusFound)

	} else {
		var check_mdp bool = CompareMDP(motdepasse, motdepasse_sql)

		if check_mdp {
			AddCookie(w, pseudo)
			check.Check_connexion = true
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			check.Check_mdp_connection = true
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

		check_pokemon_sql := SearchPokemonSQL(pokemon_name)

		data = Pokemon{pokemon_name, pokemon_types, pokemon_image}
		booster.Booster = append(booster.Booster, data)

		// quelques pokemons n'ont pas d'image sur l'API
		if pokemon_image == "" {
			i--
		}
		// evite les doublons sur la base
		if check_pokemon_sql == pokemon_name {
			continue
		} else {
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

	cookie_pseudo := ReadCookie(w, r)
	nb_pokemon := SearchNbPokemonSQL(cookie_pseudo)

	type Data struct {
		Jeux       Jeux_slice
		Nb_pokemon int
	}

	data := Data{
		Jeux:       *jeux,
		Nb_pokemon: nb_pokemon,
	}

	tmpl.Execute(w, data)
}
