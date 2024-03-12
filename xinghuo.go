package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type XHResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

func getXinghuo(message string) string {
	URL := "http://43.138.233.250:6666/GPT/ask"
	query := url.Values{}
	query.Add("question", message)
	urlWithParams := URL + "?" + query.Encode()
	errStr := "出现错误"
	request, err := http.NewRequest("GET", urlWithParams, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
		return errStr
	}

	res, err := (&http.Client{}).Do(request)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
		return errStr
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatalf("HTTP request failed with status code: %d", res.StatusCode)
		return errStr
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return errStr
	}

	resp := &XHResp{}
	if err := json.Unmarshal(body, resp); err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
		return errStr
	}

	log.Println("Response:", resp)
	return resp.Data
}
