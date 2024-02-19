package api

import (
	"bytes"
	"comfyui/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	token          = getEnv("token")
	chatURL        = "https://api.openai.com/v1/chat/completions"
	defaultContent = "이미지를 생성해주는 AI에 요청할 수 있는 적절한 입력 프롬프트 3개를 아래 키워드를 이용해 만들어 주세요."
)

func CallGPT(input string) ([]string, error) {
	m := make(map[string]interface{})
	m["model"] = "gpt-3.5-turbo"
	messages := make([]map[string]interface{}, 0)
	messages = append(messages, map[string]interface{}{
		"role":    "user",
		"content": defaultContent + "\n" + input,
	})
	m["messages"] = messages

	data, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(data)
	req, _ := http.NewRequest("POST", chatURL, buf)
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response ", err)
		return nil, err
	}

	var rsp types.GPTResponse
	if err := json.Unmarshal(body, &rsp); err != nil {
		return nil, err
	}

	var output []string
	for _, choice := range rsp.Choices {
		tr := strings.Split(choice.Message.Content, "\n")
		for _, split := range tr {
			output = append(output, split)
		}
	}

	return output, nil
}
