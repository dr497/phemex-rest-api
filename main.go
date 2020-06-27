package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := PhemexClient{Client: &http.Client{}, Api: "", Secret: []byte("")}
	resp, _ := client.getPosition("BTC")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
