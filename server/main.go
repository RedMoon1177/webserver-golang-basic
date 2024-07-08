package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/url":
		url(w, r)
	case "/body":
		body(w, r)
	default:
		w.Write([]byte("Hello World"))
	}
}

func url(w http.ResponseWriter, r *http.Request) {
	var name string

	if r.Method == http.MethodPost {
		r.ParseForm()
		name = r.FormValue("name")
	} else {
		name = r.URL.Query().Get("name")
	}

	if name == "" {
		name = "Guest"
	}
	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func body(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read body", http.StatusBadRequest)
		return
	}

	var data map[string]string
	if err := json.Unmarshal(body, &data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	name, ok := data["name"]
	if !ok {
		name = "Guest"
	}

	w.Write([]byte(fmt.Sprintf("Hello, %s!", name)))
}

func main() {
	http.HandleFunc("/", handleFunc)
	http.ListenAndServe("", nil)
}
