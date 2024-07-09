package agent

import (
	"fmt"
	"strings"
)

type Agent struct{}

func NewAgent() *Agent {
	return &Agent{}
}

func (a *Agent) CreatePrompt(tools []Tool) {
	var toolStrings []string
	var toolNames []string
	for _, tool := range tools {
		toolStrings = append(toolStrings, fmt.Sprintf("%s: %s", tool.name, tool.description))
		toolNames = append(toolNames, tool.name)
	}
	toolStringsJoined := strings.Join(toolStrings, "\n")
	toolNamesJoined := strings.Join(toolNames, ",")
	formatInstructions := strings.Replace(FORMAT_INSTRUCTIONS, "{{.ToolNames}}", toolNamesJoined, 1)
	fmt.Println(toolStringsJoined)
	fmt.Println(formatInstructions)

}
