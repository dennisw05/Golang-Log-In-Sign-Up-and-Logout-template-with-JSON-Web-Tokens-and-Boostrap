package handlers

import (
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseGlob("./templates/*")
	if err != nil {
		http.Error(w, "Error parsing templates: "+err.Error(), http.StatusInternalServerError)
		return
	}
	username, ok := r.Context().Value("username").(string)
	if !ok {
		username = "Guest"
	}
	err = tpl.ExecuteTemplate(w, "home.gohtml", username)
	if err != nil {
		http.Error(w, "Error executing template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
