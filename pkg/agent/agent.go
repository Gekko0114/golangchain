package agent

import (
	"fmt"
	"golangchain/pkg/lib"
	"golangchain/pkg/parser"
	"golangchain/pkg/prompt"
	"strings"
)

type Agent struct {
	LLMChain  *lib.Pipeline
	Prompt    *prompt.ChatPromptTemplate
	UserInput string
}

func NewAgent(tools map[string]Tool, llm lib.Runnable) (*Agent, error) {
	parser := parser.NewStrOutputParser()
	prompt, err := CreatePrompt(tools)
	if err != nil {
		return nil, err
	}
	pl := lib.NewPipeline()
	pl.Pipe(prompt).Pipe(llm).Pipe(parser)
	agent := &Agent{
		LLMChain: pl,
		Prompt:   prompt,
	}

	return agent, nil
}

func CreatePrompt(tools map[string]Tool) (*prompt.ChatPromptTemplate, error) {
	var toolStrings []string
	var toolNames []string
	for _, tool := range tools {
		toolStrings = append(toolStrings, fmt.Sprintf("%s: %s", tool.name(), tool.description()))
		toolNames = append(toolNames, tool.name())
	}
	toolStringsJoined := strings.Join(toolStrings, "\n")
	toolNamesJoined := strings.Join(toolNames, ",")
	formatInstructions := strings.Replace(FORMAT_INSTRUCTIONS, "{{.ToolNames}}", toolNamesJoined, 1)
	instructions := []string{
		SYSTEM_MESSAGE_PREFIX,
		toolStringsJoined,
		formatInstructions,
		SYSTEM_MESSAGE_SUFFIX,
	}
	instruction := strings.Join(instructions, "\n\n")
	prompt, err := prompt.NewChatPromptTemplate(instruction, HUMAN_MESSAGE)
	if err != nil {
		return nil, err
	}
	return prompt, nil
}

func (a *Agent) Plan(intermediateSteps []string) (any, error) {
	agentScratchpad := strings.Join(intermediateSteps, "\n")
	m := map[string]string{
		"Input":            a.UserInput,
		"Agent_scratchpad": agentScratchpad,
	}
	prompt, err := a.Prompt.Invoke(m)
	if err != nil {
		return nil, err
	}
	agentDecision, err := a.LLMChain.Invoke(prompt)
	if err != nil {
		return nil, err
	}
	fmt.Println(agentDecision)
	return agentDecision, err
}
