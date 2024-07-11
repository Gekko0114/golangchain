package agent

import (
	"fmt"
	"golangchain/pkg/lib"
)

type AgentExecutor struct {
	Agent         *Agent
	Tools         map[string]Tool
	MaxIterations int
}

func InitializeAgent(tools map[string]Tool, llm lib.Runnable) (*AgentExecutor, error) {
	agent, err := NewAgent(tools, llm)
	if err != nil {
		return nil, err
	}

	return &AgentExecutor{
		Agent:         agent,
		Tools:         tools,
		MaxIterations: 15,
	}, nil
}

func (a *AgentExecutor) Invoke(input any) (any, error) {
	switch input.(type) {
	case string:
		a.Agent.UserInput = input.(string)
	default:
		return nil, nil
	}
	output, err := a.call()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (a *AgentExecutor) takeNextStep(intermediateSteps []string) (any, error) {
	var results []any
	output, err := a.Agent.Plan(intermediateSteps)
	if err != nil {
		return nil, err
	}

	action := output.(string)
	var observation any
	if tool, ok := a.Tools[action]; ok {
		observation, err = tool.run("aaaaa")
		if err != nil {
			return nil, err
		}
	}
	results = append(results, observation)

	return results, nil
}

func (a *AgentExecutor) call() (any, error) {
	iterations := 0
	var intermediateSteps []string
	var nextStepOutput any
	for a.MaxIterations > iterations {
		nextStepOutput, err := a.takeNextStep(intermediateSteps)
		if err != nil {
			return nil, err
		}
		// TODO: fix
		fmt.Println(nextStepOutput)
		iterations += 1
	}
	return nextStepOutput, nil
}
