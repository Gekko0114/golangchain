package openai

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestSendMessage(t *testing.T) {
	tests := []struct {
		name     string
		message  []Message
		expected *Response
	}{
		{
			name: "normal condition",
			message: []Message{
				{Role: "system", Content: "mock test server"},
				{Role: "user", Content: "mock test question"},
			},
			expected: &Response{
				Choices: []Choice{
					{Message: Message{Role: "assistant", Content: "This is the test response"}},
				},
			},
		},
	}
	// mock
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", openaiURL,
		func(req *http.Request) (*http.Response, error) {

			response := Response{Choices: []Choice{
				{Message: Message{Role: "assistant", Content: "This is the test response"}},
			},
			}
			resp, err := httpmock.NewJsonResponse(200, response)
			if err != nil {
				return httpmock.NewStringResponse(500, ""), nil
			}
			return resp, nil
		})

	client, _ := NewChatOpenAI("gpt-3.5-turbo")
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			have, err := client.SendMessage(tc.message)
			if err != nil {
				t.Fatalf("Error happens: %v", err)
			}
			if !reflect.DeepEqual(have, tc.expected) {
				t.Fatalf("unexpected entries: %v != %v", have, tc.expected)
			}
		})
	}
}
