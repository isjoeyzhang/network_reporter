package main

import (
	"fmt"
	"net/http"
	"strings"
)

func request(url string, b string) {

	method := "GET"

	payload := strings.NewReader(b)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

}
