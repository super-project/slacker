package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var url = "http://xkcd.com/info.0.json"

type XkcdResponse struct {
	Xkcd string `json:"img"`
}

func Xkcd() string {
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	var j XkcdResponse
	json.Unmarshal(body, &j)
	return j.Xkcd
}
