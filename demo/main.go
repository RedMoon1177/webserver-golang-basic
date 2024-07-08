package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func login(w http.ResponseWriter, r *http.Request) {
	var fileName = "login.html"
	t, err := template.ParseFiles(fileName)
	if err != nil {
		fmt.Println("Error when parsing file", err)
		return
	}
	err = t.ExecuteTemplate(w, fileName, nil)
	if err != nil {
		fmt.Println("Error when executing template", err)
		return
	}
}

var userDB = map[string]string{
	"Wallace": "goodPassword",
}

func loginSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if userDB[username] == password {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "You've now logged in. Welcome to GoLang Dojo!")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Didn't find matching credentials!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/login": // todo: input boxes & button for credentials
		login(w, r)
	case "/login-submit": // todo: handle login credentials
		loginSubmit(w, r)
	default:
		fmt.Fprintf(w, "Sup Ninjas")
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("", nil)
}
