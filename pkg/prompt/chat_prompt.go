package prompt

import (
	"bytes"
	"fmt"
	"golangchain/pkg/llm"
	"text/template"
)

type ChatTemplate struct {
	system *template.Template
	human  *template.Template
}

type ChatPromptTemplate struct {
	template *ChatTemplate
}

func NewChatPromptTemplate(system string, human string) (*ChatPromptTemplate, error) {
	sysTmpl, err := template.New("system").Parse(system)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}
	humanTmpl, err := template.New("human").Parse(human)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %w", err)
	}
	Chat := &ChatTemplate{
		system: sysTmpl,
		human:  humanTmpl,
	}
	return &ChatPromptTemplate{
		template: Chat,
	}, nil
}

func (t *ChatPromptTemplate) Invoke(input any) (any, error) {
	var sbuf bytes.Buffer
	if err := t.template.system.Execute(&sbuf, input); err != nil {
		return "", err
	}
	sMes := llm.Message{
		Role:    "system",
		Content: sbuf.String(),
	}
	var hbuf bytes.Buffer
	if err := t.template.human.Execute(&hbuf, input); err != nil {
		return "", err
	}
	hMes := llm.Message{
		Role:    "user",
		Content: hbuf.String(),
	}
	messages := []llm.Message{
		sMes,
		hMes,
	}
	return messages, nil
}
