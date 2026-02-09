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
	Motdepasse TEXT  NOT NULL,
	Email TEXT NOT NULL UNIQUE

		);`

	_, err = db.Exec(create_user_table)
	if err != nil {
		log.Fatal(err)
	}
	create__jeux_table :=
		` CREATE TABLE if not exists jeux
		(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Types TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		FOREIGN KEY (user_id) REFERENCES user(id)
		);

		`

	_, err = db.Exec(create__jeux_table)
	if err != nil {
		log.Fatal(err)
	}
}
func TableUser(Pseudo *string, Motdepasse *string, Email *string) {
	db, err := sql.Open("sqlite3", "./donnerprojet.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO user (Username , Motdepasse ,Email) 
     VALUES(?, ?, ?);`, Pseudo, Motdepasse, Email)
	if err != nil {
		log.Fatal(err)
	}
}
