package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/credkellar-boop/Mon-XDR/pkg/schema"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type Analyzer struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

// NewAnalyzer initializes the Gemini client. Ensure GEMINI_API_KEY is in your environment variables.
func NewAnalyzer(ctx context.Context) (*Analyzer, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return nil, fmt.Errorf("GEMINI_API_KEY environment variable is not set")
	}

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	// Using gemini-1.5-pro for complex cross-domain reasoning and context windows
	model := client.GenerativeModel("gemini-1.5-pro")
	
	// Enforce JSON output for programmatic handling
	model.ResponseMIMEType = "application/json"

	return &Analyzer{
		client: client,
		model:  model,
	}, nil
}

// AnalyzeTelemetry sends the payload to Gemini and returns a structured threat assessment.
func (a *Analyzer) AnalyzeTelemetry(ctx context.Context, payload schema.TelemetryPayload) (*schema.ThreatAnalysisResult, error) {
	prompt := fmt.Sprintf(`Analyze the following XDR telemetry for zero-day threats or anomalies. 
	Return a JSON object matching this structure exactly: {"threat_level": float, "is_zero_day": boolean, "explanation": "string"}.
	Telemetry Data: %+v`, payload)

	resp, err := a.model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return nil, fmt.Errorf("gemini analysis failed: %w", err)
	}

	if len(resp.Candidates) == 0 || len(resp.Candidates[0].Content.Parts) == 0 {
		return nil, fmt.Errorf("no response generated")
	}

	var result schema.ThreatAnalysisResult
	part := resp.Candidates[0].Content.Parts[0]

	if textPart, ok := part.(genai.Text); ok {
		if err := json.Unmarshal([]byte(textPart), &result); err != nil {
			return nil, fmt.Errorf("failed to parse AI JSON response: %w", err)
		}
	} else {
		return nil, fmt.Errorf("unexpected response type from API")
	}

	result.EventID = payload.EventID
	return &result, nil
}

// Close cleans up the client connection.
func (a *Analyzer) Close() {
	a.client.Close()
}
