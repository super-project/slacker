package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var url = "http://tambal.azurewebsites.net/joke/random"

type JokeResponse struct {
	Joke string `json:"joke"`
}

func Joke() string {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var j JokeResponse
	json.Unmarshal(body, &j)
	return j.Joke
}
