package handlers

import (
	"loginRegisterTemplate/helpers"
	"loginRegisterTemplate/repositories"
	"loginRegisterTemplate/types"
	"net/http"
	"text/template"

	"gorm.io/gorm"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("./templates/*")
	if err != nil {
		http.Error(w, "Error parsing templates: "+err.Error(), http.StatusInternalServerError)
		return
	}
	loginErrorMessage := helpers.GetErrorCookie("login-error", w, r)
	err = tpl.ExecuteTemplate(w, "login.gohtml", loginErrorMessage)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	user, err := repositories.GetUser(types.User{Username: username})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			helpers.ErrorCookie("login-error", "Incorrect login credentials 1!", w)
		} else {
			helpers.ErrorCookie("login-error", "Something went wrong! Please try again.", w)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if !helpers.CheckPasswords(password, user.Password) {
		helpers.ErrorCookie("login-error", "Incorrect login credentials! 2", w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	err = helpers.SetJWTToken(user.Username, w)
	if err != nil {
		helpers.ErrorCookie("login-error", "Something went wrong! Please try again.", w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
