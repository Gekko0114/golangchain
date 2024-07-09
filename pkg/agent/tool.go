package agent

var AvailableTools = map[string]Tool{
	"llm-math": {
		name:        "Calculator",
		description: "Useful for when you need to answer questions about math.",
	},
	"serpapi": {
		name:        "Search",
		description: "A search engine. Useful for when you need to answer questions about current events. Input should be a search query.",
	}}

type Tool struct {
	name        string
	description string
}

func LoadTools(tools []string) []Tool {
	confirmedTools := []Tool{}
	for _, toolname := range tools {
		if tool, ok := AvailableTools[toolname]; ok {
			confirmedTool := tool
			confirmedTools = append(confirmedTools, confirmedTool)
		}
	}
	return confirmedTools
}
