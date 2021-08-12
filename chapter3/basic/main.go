package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {

	// GET

	resp, err := http.Get("https://github.com/")
	if err != nil {
		log.Println("Error in HTTP get")
	}

	// Print HTTP StatusCode
	fmt.Println(resp.StatusCode)

	// Read and display response body

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Panicln("Error in reading response bodys")
	}

	fmt.Println(string(body))

	// closing reponse body

	defer resp.Body.Close()

	// HEAD

	resp, err = http.Head("https://github.com/")

	if err != nil {
		log.Panicln("Error in reading response bodys")
	}

	defer resp.Body.Close()

	// Status
	fmt.Println(resp.Status)

	form := url.Values{}

	form.Add("name", "yuvan11")

	// POST

	resp, err = http.Post(
		"https://github.com/",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)

	if err != nil {
		log.Panicln(err)
	}

	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp)

	req, err := http.NewRequest("DELETE", "https://www.google.com/robots.txt", nil)
	if err != nil {
		log.Panicln(err)
	}
	var client http.Client
	resp, err = client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	req, _ = http.NewRequest("PUT", "https://www.google.com/robots.txt", strings.NewReader(form.Encode()))
	resp, err = client.Do(req)
	if err != nil {
		log.Panicln(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
}
