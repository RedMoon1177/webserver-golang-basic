package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	// Plain old URL - http://localhost
	response, err := http.Get("http://localhost")
	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// URL key-value form - http://localhost/url?name=Wallace
	response, err = http.PostForm(
		"http://localhost/url",
		url.Values{"name": {"Wallace"}},
	)

	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}

	// http://localhost/body
	// with body in json: {"name": "Wallace"}
	type Ninja struct {
		Name string `json:"name"`
	}

	wallace := Ninja{Name: "Wallace"}
	wallaceJson, _ := json.Marshal(wallace)

	response, err = http.Post(
		"http://localhost/body",
		"application/json",
		bytes.NewBuffer(wallaceJson),
	)

	if err != nil {
		fmt.Println(err)
	} else {
		data, _ := io.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
