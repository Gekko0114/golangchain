package llm

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

type ChatOpenAI struct {
	apiKey     string
	model      string
	httpClient *http.Client
}

func NewChatOpenAI(model string) (*ChatOpenAI, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}
	return &ChatOpenAI{
		apiKey,
		model,
		&http.Client{},
	}, nil
}

func (c *ChatOpenAI) Invoke(input any) (any, error) {
	var requestBody Request
	switch input.(type) {
	case []Message:
		messages, _ := input.([]Message)
		requestBody = Request{
			Model:    c.model,
			Messages: messages,
		}
	case string:
		content, _ := input.(string)
		requestBody = Request{
			Model: c.model,
			Messages: []Message{
				{
					Role:    "user",
					Content: content,
				},
			},
		}
	default:
		return nil, fmt.Errorf("Error while calling chatOpenAI invoke")
	}

	requestBodyJSON, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling request body: %w", err)
	}
	// Create a new HTTP request
	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, fmt.Errorf("Error creating HTTP request: %w", err)
	}

	// Set the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending HTTP request: %w", err)
	}
	defer resp.Body.Close()
	// Read the response body
	respBody, err := io.ReadAll(resp.Body)

	// Unmarshal the response body to the Response struct
	var response Response
	err = json.Unmarshal(respBody, &response)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling response body: %w", err)
	}
	return &response, nil
}
