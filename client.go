package main

import (
	"io/ioutil"
	"net/http"
)

var apiURL = "https://configd.gophersauce.com/get_configuration"

func ReadBody(r *http.Response) string {
	body, _ := ioutil.ReadAll(r.Body)
	return string(body)
}
