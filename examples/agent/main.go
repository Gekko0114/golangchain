package main

import (
	"fmt"
	"golangchain/pkg/agent"
	"golangchain/pkg/llm"
)

func main() {
	llm, err := llm.NewChatOpenAI("gpt-3.5-turbo")
	if err != nil {
		fmt.Println(err)
	}
	tools := agent.LoadTools([]string{"serpapi"})
	agentExecutor, err := agent.InitializeAgent(tools, llm)
	if err != nil {
		fmt.Println(err)
	}
	prompt := "青を英語で言うと、何か？"
	result, err := agentExecutor.Invoke(prompt)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Response: %+v\n", result)
}
