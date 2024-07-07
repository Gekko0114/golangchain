package main

import (
	"fmt"
	"golangchain/pkg/openai"
)

func main() {
	llm, err := openai.NewChatOpenAI("gpt-3.5-turbo")
	if err != nil {
		fmt.Println(err)
	}

	msg := []openai.Message{
		{Role: "system", Content: ""},
		{Role: "user", Content: "こんにちは"},
	}

	response, err := llm.SendMessage(msg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Response:", response.Choices[0].Message.Content)

}
