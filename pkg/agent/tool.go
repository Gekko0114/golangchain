package agent

import (
	"golang.org/x/exp/slices"
)

var AvailableTools = []string{"llm-math", "serpapi"}

type Tool struct {
	name string
}

func LoadTools(tools []string) []Tool {
	confirmedTools := []Tool{}
	for _, tool := range tools {
		if slices.Contains(AvailableTools, tool) {
			confirmedTool := &Tool{name: tool}
			confirmedTools = append(confirmedTools, *confirmedTool)
		}
	}
	return confirmedTools
}
