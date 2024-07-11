package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var AvailableTools = map[string]Tool{
	"serpapi": &SerpAPI{}}

type Tool interface {
	name() string
	description() string
	run(input any) (any, error)
}

func LoadTools(tools []string) map[string]Tool {
	confirmedTools := map[string]Tool{}
	for _, toolname := range tools {
		if tool, ok := AvailableTools[toolname]; ok {
			confirmedTool := tool
			confirmedTools[toolname] = confirmedTool
		}
	}
	return confirmedTools
}

type OrganicResult struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

type SerpApiResponse struct {
	OrganicResults []OrganicResult `json:"organic_results"`
}

type SerpAPI struct {
}

func (s *SerpAPI) name() string {
	return "Search"
}

func (s *SerpAPI) description() string {
	return "A search engine. Useful for when you need to answer questions about current events. Input should be a search query."
}

func (s *SerpAPI) run(input any) (any, error) {
	apiKey := os.Getenv("SERPAPI_API_KEY")
	searchQuery := input.(string)
	serpApiURL := "https://serpapi.com/search"

	params := url.Values{}
	params.Add("engine", "google")
	params.Add("q", searchQuery)
	params.Add("api_key", apiKey)

	resp, err := http.Get(fmt.Sprintf("%s?%s", serpApiURL, params.Encode()))
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response SerpApiResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}
	return response, nil
}
