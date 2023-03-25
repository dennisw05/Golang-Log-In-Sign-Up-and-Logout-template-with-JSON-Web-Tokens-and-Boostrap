package handlers

import (
	"log"
	"loginRegisterTemplate/helpers"
	"loginRegisterTemplate/repositories"
	"loginRegisterTemplate/types"
	"net/http"
	"text/template"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("./templates/*")
	if err != nil {
		http.Error(w, "Error parsing templates: "+err.Error(), http.StatusInternalServerError)
		return
	}
	signupErrorMessage := helpers.GetErrorCookie("signup-error", w, r)
	err = tpl.ExecuteTemplate(w, "signup.gohtml", signupErrorMessage)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}

func SignupPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	passwordConfirm := r.Form.Get("password-confirm")
	if !helpers.CheckSignupUsername(username, w) {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	hashedPassword, ok := helpers.CheckSignupPassword(password, passwordConfirm, w)
	if !ok {
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
	}
	user := types.User{Username: username, Password: hashedPassword}
	err := repositories.CreateUser(user)
	if err != nil {
		log.Println("Could not create the user!")
		helpers.ErrorCookie("signup-error", "Something went wrong! Please try again.", w)
		http.Redirect(w, r, "/signup", http.StatusSeeOther)
		return
	}
	err = helpers.SetJWTToken(user.Username, w)
	if err != nil {
		helpers.ErrorCookie("login-error", "Could'nt log in automatically! Please log in.", w)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
