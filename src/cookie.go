package pokemon

import "net/http"

func AddCookie(w http.ResponseWriter, r *http.Request) {

	///////////////////////////////////////////////////////
	// REQUETE SQL POUR RECUP ID						 //
	///////////////////////////////////////////////////////

	cookie := &http.Cookie{
		Name: "cookie",
		// Value: id,
		Path: "/",
	}

	http.SetCookie(w, cookie)
}
