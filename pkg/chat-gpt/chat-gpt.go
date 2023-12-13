package chat_gpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

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
	payload := map[string]interface{}{
		"model":      "gpt-3.5-turbo",
		"messages":   []interface{}{map[string]interface{}{"role": "system", "content": question}},
		"max_tokens": 2048,
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
		log.Fatalf("Error sending HTTP request: %v", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
		return "", err
	}
	// Unmarshal JSON response
	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalf("Error decoding JSON response: %v", err)
		return "", err
	}
	// TODO NEED TO CATCH ERRORS HERE
	// Extract the choices array from the JSON response
	choices, ok := data["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", errors.New("malformed response: missing or empty choices array")
	}

	// Extract the first choice from the array
	firstChoice, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", errors.New("malformed response: invalid format for the first choice")
	}

	// Extract the message object from the first choice
	message, ok := firstChoice["message"].(map[string]interface{})
	if !ok {
		return "", errors.New("malformed response: missing or invalid message field")
	}

	// Extract the content string from the message object
	content, ok := message["content"].(string)
	if !ok {
		return "", errors.New("malformed response: missing or invalid content field")
	}

	return content, nil
}
