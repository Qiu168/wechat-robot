package main

import (
	"io"
	"net/http"
	"strings"
)

func doRpc(body, url, method string) string {
	client := &http.Client{}
	request, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return ""
	}
	resp, err := client.Do(request)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	message, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(message)
}
