package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
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
	request, _ := http.NewRequest("GET", urlWithParams, nil)
	fmt.Println(request.URL.Path)
	//method := "GET"
	res, _ := (&http.Client{}).Do(request)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	ans := string(body)
	fmt.Println("data=" + ans)
	resp := &XHResp{}
	decoder := json.NewDecoder(strings.NewReader(ans))
	err = decoder.Decode(resp)
	if err != nil {
		fmt.Errorf("err:%v", err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Data)
	return resp.Data
}
