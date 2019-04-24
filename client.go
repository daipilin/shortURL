package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	create("http://baidu.com/adfadfa")
	query("10001")
}

func create(longUrl string) {
	response, error := http.PostForm("http://localhost:8000/create",
		url.Values{
			"url": {longUrl},
		})
	if error != nil {
		fmt.Println("failed")
	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("failed")
	}
	fmt.Println(string(body))
}

func query(shortUrl string) {
	response, error := http.PostForm("http://localhost:8000/query",
		url.Values{
			"shortUrl": {shortUrl},
		})
	if error != nil {
		fmt.Println("failed")
	}
	defer response.Body.Close()
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		fmt.Println("failed")
	}
	fmt.Println(string(body))
}