package helpers

import (
	"net/http"
	"time"
)

func ErrorCookie(cookieName, value string, w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     cookieName,
		Value:    value,
		Expires:  time.Now().Add(time.Minute * 1),
		HttpOnly: true,
	})
}

func GetErrorCookie(cookieName string, w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		return ""
	}
	http.SetCookie(w, &http.Cookie{
		Name:   cookieName,
		Value:  "",
		MaxAge: -1,
	})
	return cookie.Value
}
