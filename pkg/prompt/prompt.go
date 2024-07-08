package prompt

import (
	"bytes"
	"fmt"
	"text/template"
)

type PromptTemplate struct {
	template *template.Template
}

func NewPromptTemplate(input string) (*PromptTemplate, error) {
	tmpl, err := template.New("tmpl").Parse(input)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}

	return &PromptTemplate{
		template: tmpl,
	}, nil
}

func (t *PromptTemplate) Invoke(input any) (string, error) {
	var buf bytes.Buffer
	if err := t.template.Execute(&buf, input); err != nil {
		return "", err
	}
	return buf.String(), nil
}
