package prompt

import (
	"testing"
)

func TestInvoke(t *testing.T) {
	tests := []struct {
		name     string
		template string
		input    any
		expected string
	}{
		{
			name:     "template is string",
			template: "template",
			expected: "template",
		},
		{
			name:     "template includes input",
			template: "the meaning of {{.Word}}?",
			input: map[string]string{
				"Word": "satisfaction",
			},
			expected: "the meaning of satisfaction?",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			prompt, err := NewPromptTemplate(tc.template)
			if err != nil {
				t.Fatalf("Error happens: %v", err)
			}
			have, err := prompt.Invoke(tc.input)
			if err != nil {
				t.Fatalf("Error happens: %v", err)
			}
			if have != tc.expected {
				t.Fatalf("unexpected string: %v != %v", have, tc.expected)
			}
		})
	}
}
