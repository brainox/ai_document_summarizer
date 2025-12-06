package external_models

import (
	"ai_document_summarizer/external/thirdparty/openrouter"
)

type OpenRouterClient struct {
	client *openrouter.Client
}

func NewOpenRouterClient(apiKey, model string) *OpenRouterClient {
	return &OpenRouterClient{
		client: openrouter.NewClient(apiKey, model),
	}
}

func (o *OpenRouterClient) AnalyzeDocument(text string) (string, error) {
	return o.client.Analyze(text)
}
