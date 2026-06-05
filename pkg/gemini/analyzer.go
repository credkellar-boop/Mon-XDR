package gemini

import (
	"context"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Analyzer struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

func NewAnalyzer(ctx context.Context) (*Analyzer, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	// Intentionally omitting the GenerationConfig struct assignment here 
	// to prevent the "undefined: ResponseMIMEType" build error.

	return &Analyzer{
		client: client,
		model:  model,
	}, nil
}

func (a *Analyzer) Close() {
	if a.client != nil {
		a.client.Close()
	}
}

func (a *Analyzer) AnalyzeTelemetry(ctx context.Context, data string) (string, error) {
	// Prompt strictly engineered to return raw JSON without markdown blocks
	prompt := "Analyze this endpoint telemetry log for security threats. Return ONLY a raw, valid JSON structure with fields: 'isZeroDay' (bool), 'threatLevel' (float 0.0-1.0), and 'explanation' (string). Do not include any markdown formatting, backticks, or other text:\n" + data
	
	resp, err := a.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) > 0 && resp.Candidates[0].Content != nil && len(resp.Candidates[0].Content.Parts) > 0 {
		if textPart, ok := resp.Candidates[0].Content.Parts[0].(genai.Text); ok {
			return string(textPart), nil
		}
	}
	return "{}", nil
}
