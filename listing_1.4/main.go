package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://example.com"
	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	resp.Body.Close()
}
