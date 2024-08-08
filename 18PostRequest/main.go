package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web get request")
	// PerfomPostJsonRequest()
	PerformPostFormRequest()
}

func PerfomPostJsonRequest() {
	const myUrl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
		"Coursename":"lets go with go lang",
		"price":"0",
		"platform":"LCD.in"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	//formdata

	data := url.Values{}
	data.Add("firstname", "majid")
	data.Add("lastname", "beg")
	data.Add("email", "m@bgo.dev")

	response, err := http.PostForm(myUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
