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
	a.Agent.UserInput = input.(string)
	output, err := a.call()
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (a *AgentExecutor) takeNextStep(intermediateSteps []string) (string, string, error) {
	var observation string
	nextaction, finalanswer, err := a.Agent.Plan(intermediateSteps)
	if err != nil {
		return "", "", err
	}
	if len(finalanswer) > 1 {
		return "", finalanswer, nil
	}
	if tool, ok := a.Tools[nextaction.Action.Action_name]; ok {
		observation, err = tool.run(nextaction.Action.Action_input)
		if err != nil {
			return "", "", fmt.Errorf("error during takeNextStep: %w", err)
		}
	}
	nextactionstring := fmt.Sprintf("Thought: %s\nAction_name: %s\nAction_input: %s", nextaction.Thought, nextaction.Action.Action_name, nextaction.Action.Action_input)

	return observation, nextactionstring, nil
}

func (a *AgentExecutor) call() (any, error) {
	iterations := 0
	var nextStepOutput any
	var intermediateSteps []string
	for a.MaxIterations > iterations {
		nextStepOutput, pastaction, err := a.takeNextStep(intermediateSteps)
		if err != nil {
			return nil, err
		}
		if nextStepOutput == "" {
			return pastaction, nil
		}
		jsonData, err := json.Marshal(nextStepOutput)
		if err != nil {
			return nil, err
		}
		intermediateSteps = append(intermediateSteps, pastaction)
		intermediateSteps = append(intermediateSteps, string(jsonData))
		iterations += 1
	}
	return nextStepOutput, nil
}
