package agent

import (
	"encoding/json"
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
		MaxIterations: 5,
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
	var observation any
	output, err := a.Agent.Plan(intermediateSteps)
	if err != nil {
		return nil, err
	}
	if tool, ok := a.Tools[output.Action.Action_name]; ok {
		observation, err = tool.run(output.Action.Action_input)
		if err != nil {
			return nil, fmt.Errorf("error during takeNextStep: %w", err)
		}
	}

	return observation, nil
}

func (a *AgentExecutor) call() (any, error) {
	iterations := 0
	var nextStepOutput any
	var intermediateSteps []string
	for a.MaxIterations > iterations {
		nextStepOutput, err := a.takeNextStep(intermediateSteps)
		if err != nil {
			return nil, err
		}
		jsonData, err := json.Marshal(nextStepOutput)
		if err != nil {
			return nil, err
		}
		intermediateSteps = append(intermediateSteps, string(jsonData))
		iterations += 1
	}
	return nextStepOutput, nil
}
