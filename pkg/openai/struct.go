package openai

// message represents a single message in the chat
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Request represents the OpenAI API request payload
type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Response represents the OpenAI API response
type Response struct {
	Choices []Choice `json:"choices"`
}

// Choice represents a single choice in the OpenAI API response
type Choice struct {
	Message Message `json:"message"`
}
