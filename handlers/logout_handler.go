package handlers

import (
	"net/http"
	"time"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	expireTime := time.Now().Add(-time.Hour)
	jwtCookie := &http.Cookie{
		Name:     "jwt-token",
		Value:    "",
		Expires:  expireTime,
		HttpOnly: true,
	}
	http.SetCookie(w, jwtCookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
