package pokemon

import (
	"log"
	"net/http"
)

func AddCookie(w http.ResponseWriter, pseudo string) {
	cookie := &http.Cookie{
		Name:  "Session",
		Value: pseudo,
		Path:  "/",
	}

	http.SetCookie(w, cookie)
}

func ReadCookie(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("Session")

	if err != nil {
		log.Fatal(err)
	}

	pseudo_cookie := cookie.Value
	return pseudo_cookie
}

func DeleteCookie(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "Session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}

	http.SetCookie(w, cookie)
}
