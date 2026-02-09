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
		`CREATE TABLE IF NOT EXISTS user (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Username TEXT NOT NULL UNIQUE,
			Email TEXT NOT NULL UNIQUE,
			Motdepasse TEXT  NOT NULL
		);`

	_, err = db.Exec(create_user_table)
	if err != nil {
		log.Fatal(err)
	}

	create__jeux_table :=
		`CREATE TABLE IF NOT EXISTS jeux (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			Name TEXT NOT NULL,
			Types TEXT NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user(id)
		);`

	_, err = db.Exec(create__jeux_table)
	if err != nil {
		log.Fatal(err)
	}
}
func AddUser(pseudo string, email string, motdepasse string) {
	db, err := sql.Open("sqlite3", "./donnerprojet.sqlite")
	ReadError(err)
	defer db.Close()

	_, err = db.Exec(
		`INSERT INTO user (Username, Email , Motdepasse) 
     	 VALUES(?, ?, ?);`, pseudo, email, motdepasse)

	ReadError(err)
}
