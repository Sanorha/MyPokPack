package pokemon

import (
	"net/http"
)

func AddCookie(w http.ResponseWriter) {

	///////////////////////////////////////////////////////
	// REQUETE SQL POUR RECUP ID						 //
	///////////////////////////////////////////////////////

	cookie := &http.Cookie{
		Name:  "cookie",
		Value: "toto",
		Path:  "/",
	}

	http.SetCookie(w, cookie)
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}
