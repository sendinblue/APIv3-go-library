package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://api.sendinblue.com/v3/contacts"

	payload := strings.NewReader("{\"updateEnabled\":false}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	
	req.Header.Add("api-key", "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")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}