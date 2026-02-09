package pokemon

import (
	"net/http"
)

func AddCookie(w http.ResponseWriter, pseudo string) {
	cookie := &http.Cookie{
		Name:  "cookie",
		Value: pseudo,
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
