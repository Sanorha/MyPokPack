package pokemon

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func CreateTableone() {
	db, err := sql.Open("sqlite3", "./donnerprojet.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	create_user_table :=
		`CREATE TABLE if not exists user (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE,
			Email TEXT NOT NULL UNIQUE,
			Motdepasse TEXT  NOT NULL
		);`

	_, err = db.Exec(create_user_table)
	if err != nil {
		log.Fatal(err)
	}

	create_jeux_table :=
		`CREATE TABLE IF NOT EXISTS jeux (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL,
			Types TEXT NOT NULL,
			Image TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user(id)
		);`

	_, err = db.Exec(create_jeux_table)
	if err != nil {
		log.Fatal(err)
	}
}
func AddUser(Pseudo string, Email string, Motdepasse string) {
	db, err := sql.Open("sqlite3", "./donnerprojet.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec(
		`INSERT INTO user (Username, Email , Motdepasse) 
     	VALUES(?, ?, ?);`, Pseudo, Email, Motdepasse)

	if err != nil {
		log.Fatal(err)
	}
}

func SearchUserSQL(pseudo string) string {
	db, err := sql.Open("sqlite3", "./donnerprojet.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var username_sql string
	var motdepasse_sql string

	err = db.QueryRow(
		`SELECT Username, Motdepasse
		FROM user
		WHERE Username = ?;`, pseudo).Scan(&username_sql, &motdepasse_sql)

	if err != nil {
		return ""
	}

	return motdepasse_sql
}

func SearchIdUserSQL(pseudo string) string {
	db, err := sql.Open("sqlite3", "./donnerprojet.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var id_sql string

	err = db.QueryRow(
		`SELECT id
		FROM user
		WHERE Username = ?;`, pseudo).Scan(&id_sql)

	if err != nil {
		return ""
	}

	return id_sql
}

func AddJeuxSQL(jeux *Jeux, pseudo_cookie string) {
	db, err := sql.Open("sqlite3", "./donnerprojet.sqlite")

	if err != nil {
		log.Fatal(err)
	}

	id_pseudo := SearchIdUserSQL(pseudo_cookie)

	defer db.Close()

	_, err = db.Exec(
		`INSERT INTO jeux (Name, Types, Image, user_id) 
     	VALUES(?, ?, ?, ?);`, jeux.Name, jeux.Types, jeux.Image, id_pseudo)

	if err != nil {
		log.Fatal(err)
	}
}

func SearchJeuxSQL() {

}

/////////////// requete base jeux ajouter image //////////
/////////////// voir max pokemon /////////////////////////
/////////////// faire fonction random/////////////////////
/////////////// https://pokeapi.co/api/v2/pokemon/1 //////
