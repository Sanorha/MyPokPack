package pokemon

import (
	"database/sql"
	"fmt"
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

func SearchSQL(pseudo string) string {
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
		log.Fatal(err)
	}

	fmt.Println("Resultat base sql : ", username_sql, motdepasse_sql)

	return motdepasse_sql
}
