package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	url := "http://stat.gibdd.ru/places/getCPoints"
	method := "POST"

	payload := strings.NewReader("{ \"cpType\": \"0\", \"date\": \"YEAR:2020\", \"regId\": \"929\" }")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
