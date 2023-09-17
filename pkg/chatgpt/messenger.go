package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	chatGPTUrl  = "https://api.openai.com/v1/chat/completions"
	model       = "gpt-3.5-turbo"
	role        = "user"
	temperature = 0.7
)

type ChatGPT struct {
	client *http.Client
}

type chatGPTTransport struct {
	apiKey string
}

func NewChatGPTClient(apiKey string) *ChatGPT {
	return &ChatGPT{
		client: &http.Client{
			Transport: &chatGPTTransport{
				apiKey: apiKey,
			},
		},
	}
}

func (t *chatGPTTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+t.apiKey)

	return http.DefaultTransport.RoundTrip(req)
}

func (oai *ChatGPT) Post(cr *ChatGPTRequest) (*ChatGPTData, error) {
	cr.Model = model
	cr.Temperature = temperature

	payloadBytes, err := json.Marshal(cr)
	if err != nil {
		return nil, errors.New("openai.com failed to encode request payload")
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", chatGPTUrl, body)
	if err != nil {
		return nil, errors.New("openai.com failed to create request")
	}

	res, err := oai.client.Do(req)
	if err != nil || res.StatusCode != http.StatusOK {
		return nil, errors.New("openai.com failed to get response")
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("openai.com failed to read response body")
	}

	var data ChatGPTData

	err = json.Unmarshal(resBody, &data)
	if err != nil {
		return nil, errors.New("openai.com failed to unmarshal response body")
	}

	return &data, nil
}
