package main

import (
	"fmt"
	"golangchain/pkg/lib"
	"golangchain/pkg/llm"
	"golangchain/pkg/parser"
	"golangchain/pkg/prompt"
)

func main() {
	llm, err := llm.NewChatOpenAI("gpt-3.5-turbo")
	if err != nil {
		fmt.Println(err)
	}
	prompt, err := prompt.NewPromptTemplate("{{.Word}}の意味を教えて。")
	parser := parser.NewStrOutputParser()

	pipeline := lib.NewPipeline()
	pipeline.Pipe(prompt).Pipe(llm).Pipe(parser)
	m := map[string]string{
		"Word": "因果応報",
	}

	response, err := pipeline.Invoke(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Response: %+v\n", response)
}
