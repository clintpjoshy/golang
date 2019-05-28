package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type response struct {
	id   int    `json:"page"`
	name string `json:"rocket_name"`
}

func callMe() {
	resp, err := http.Get("https://api.spacexdata.com/v3/rockets")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		fmt.Println(string(body))
	}
}

func main() {
	callMe()
}
