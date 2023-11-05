package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Resp struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int    `json:"created"`
	Model   string `json:"model"`
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		} `json:"message"`
		FinishReason string `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func gpt(content string) string {
	url := "https://yewu.bcwhkj.cn/api/v2.Gptliu/search"
	method := "POST"

	payload := strings.NewReader(
		`{
			"messages": [
				{
					"role": "system",
					"content": "你是一个能干的助手."
				},
				{
					"role": "user",
					"content": "` + content + `"
				}
			],
			"model": "gpt-3.5-turbo-0613",
			"max_tokens": 3000,
			"stream": false
		}`,
	)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	req.Header.Add("Authorization", "Bearer 57C1WqVYrKEG1iumKnEuMc8sARbvlmDo")
	//req.Header.Add("User-Agent", "Apifox/1.0.0 (https://apifox.com)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	ans := string(body)
	fmt.Println(ans)
	resp := &Resp{}
	decoder := json.NewDecoder(strings.NewReader(ans))
	decoder.Decode(resp)
	content = resp.Choices[0].Message.Content
	fmt.Println(content)
	return content
}
