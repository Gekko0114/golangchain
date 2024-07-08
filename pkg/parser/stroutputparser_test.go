package parser

import (
	"golangchain/pkg/openai"
	"testing"
)

func TestInvoke(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name: "template is string",
			input: &openai.Response{
				Choices: []openai.Choice{
					{Message: openai.Message{Role: "assistant", Content: "This is the test response"}},
				},
			},
			expected: "This is the test response",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			parser := NewStrOutputParser()
			have, err := parser.Invoke(tc.input)
			if err != nil {
				t.Fatalf("Error happens: %v", err)
			}
			if have != tc.expected {
				t.Fatalf("unexpected string: %v != %v", have, tc.expected)
			}
		})
	}
}