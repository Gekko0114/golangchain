package parser

import (
	"golangchain/pkg/openai"
)

type StrOutputParser struct {
}

func NewStrOutputParser() *StrOutputParser {
	return &StrOutputParser{}
}

func (p *StrOutputParser) Invoke(input any) (any, error) {
	var output string
	res, ok := input.(*openai.Response)
	if ok {
		output = res.Choices[0].Message.Content
	} else {
		return nil, nil
	}
	return output, nil
}
