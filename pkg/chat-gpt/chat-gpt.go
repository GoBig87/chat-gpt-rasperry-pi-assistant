package chat_gpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type ChatGptRequest struct {
	Model    string           `json:"model"`
	Messages []RequestMessage `json:"messages"`
	Stream   bool             `json:"stream"`
}

type RequestMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGptResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choices"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type Choice struct {
	Message      ResponseMessage `json:"message"`
	FinishReason string          `json:"finish_reason"`
	Index        int             `json:"index"`
}

type ResponseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGptClient struct {
	ApiKey      string
	OrgID       string
	ApiEndpoint string
}

func NewChatGptClient(apiKey, orgId, apiEndpoint string) *ChatGptClient {
	return &ChatGptClient{
		ApiKey:      apiKey,
		OrgID:       orgId,
		ApiEndpoint: apiEndpoint,
	}
}

func (c *ChatGptClient) PromptChatGPT(question string) (string, error) {
	// Create a new HTTP client
	client := &http.Client{}

	// Prepare the request payload
	payload := &ChatGptRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []RequestMessage{{Role: "system", Content: question}},
		Stream:   true,
	}

	// Convert the payload to JSON
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		log.Fatalf("Error encoding JSON payload: %v", err)
		return "", err
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", c.ApiEndpoint, bytes.NewBuffer(payloadJSON))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
		return "", err
	}

	// Set headers and authentication token
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.ApiKey)
	req.Header.Set("OpenAI-Organization", c.OrgID)

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending HTTP request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		return "", err
	}
	// Unmarshal JSON response
	response := &ChatGptResponse{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Printf("Error decoding JSON response: %v", err)
		var data map[string]interface{}
		err2 := json.Unmarshal(body, &data)
		if err2 != nil {
			log.Printf("Found error JSON response: %v", data)
		}
		return "", err
	}

	// Extract the first choice from the array
	var ret string
	for _, choice := range response.Choices {
		ret = fmt.Sprintf("%s %s", ret, choice.Message.Content)
	}
	return ret, nil
}
