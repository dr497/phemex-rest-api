package main

import (
	"bytes"
	"net/http"
	"strconv"
	"time"
)

var URL = "https://api.phemex.com"

func (client *PhemexClient) signRequest(method string, path string, query string, body []byte) *http.Request {
	expiry := strconv.FormatInt(time.Now().UTC().Unix()*1000+60, 10)
	signaturePayload := path + query + expiry + string(body)
	signature := client.sign(signaturePayload)
	var req *http.Request
	if query != "" {
		req, _ = http.NewRequest(method, URL+path+"?"+query, bytes.NewBuffer(body))
	} else {
		req, _ = http.NewRequest(method, URL+path, bytes.NewBuffer(body))
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-phemex-access-token", client.Api)
	req.Header.Set("x-phemex-request-signature", signature)
	req.Header.Set("x-phemex-request-expiry", expiry)
	return req
}

func (client *PhemexClient) _get(path string, query string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("GET", path, query, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *PhemexClient) _post(path string, query string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("POST", path, query, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *PhemexClient) _delete(path string, query string, body []byte) (*http.Response, error) {
	preparedRequest := client.signRequest("DELETE", path, query, body)
	resp, err := client.Client.Do(preparedRequest)
	return resp, err
}

func (client *PhemexClient) getProducts() (*http.Response, error) {
	resp, err := client._get("/v1/exchange/public/products", "", []byte(""))
	return resp, err
}

func (client *PhemexClient) cancelOrder(symbol string, orderID int64) (*http.Response, error) {
	query := "symbol=" + symbol + "&orderID=" + strconv.FormatInt(orderID, 10)
	resp, err := client._delete("/orders/cancel", query, []byte(""))
	return resp, err
}

func (client *PhemexClient) getPosition(symbol string) (*http.Response, error) {
	query := "currency=" + symbol
	resp, err := client._get("/accounts/accountPositions", query, []byte(""))
	return resp, err
}
