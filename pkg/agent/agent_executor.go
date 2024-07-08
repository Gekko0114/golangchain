package agent

import "fmt"

type AgentExecutor struct {
	Agent         Agent
	Tools         []Tool
	MaxIterations int
}

func NewAgentExecutor() *AgentExecutor {
	return &AgentExecutor{}
}

func (a *AgentExecutor) Invoke(input any) (any, error) {
	return nil, nil
}

func (a *AgentExecutor) TakeNextStep() any {
	return nil
}

func (a *AgentExecutor) call() (any, error) {
	iterations := 0
	var nextStepOutput any
	for a.MaxIterations > iterations {
		nextStepOutput := a.TakeNextStep()
		fmt.Println(nextStepOutput)
		iterations += 1
	}
	return nextStepOutput, nil
}
