package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/generative-ai-go/genai"
	"google.com/api/option"

	"github.com/credkellar-boop/Mon-XDR/pkg/schema"
)

// GeminiAnalyzer wraps the Google GenAI Client and Model configuration
type GeminiAnalyzer struct {
	client *genai.Client
	model  *genai.GenerativeModel
}

// NewAnalyzer initializes the Gemini client with the required execution options
func NewAnalyzer(ctx context.Context) (*GeminiAnalyzer, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return nil, err
	}

	// Correctly specify settings via the GenerationConfig field structure
	model := client.GenerativeModel("gemini-1.5-flash")
	model.GenerationConfig = genai.GenerationConfig{
		ResponseMIMEType: "application/json",
	}

	return &GeminiAnalyzer{
		client: client,
		model:  model,
	}, nil
}

func (ga *GeminiAnalyzer) Close() {
	if ga.client != nil {
		ga.client.Close()
	}
}

// AnalyzeTelemetry sends the event message string to Gemini for automated threat categorization
func (ga *GeminiAnalyzer) AnalyzeTelemetry(ctx context.Context, data string) (string, error) {
	prompt := "Analyze this endpoint telemetry log for security threats. Return a JSON structure with fields: 'isZeroDay' (bool), 'threatLevel' (float 0.0-1.0), and 'explanation' (string):\n" + data
	resp, err := ga.model.GenerateContent(ctx, genai.Text(prompt))
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

func main() {
	log.Println("Starting Mon-XDR Worker...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	analyzer, err := NewAnalyzer(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Gemini analyzer: %v", err)
	}
	defer analyzer.Close()

	// Simulated incoming message channel
	msgs := make(chan []byte)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg := <-msgs:
				processMessage(ctx, analyzer, msg)
			}
		}
	}()

	<-sigChan
	log.Println("Shutting down worker gracefully...")
}

func processMessage(ctx context.Context, analyzer *GeminiAnalyzer, msg []byte) {
	var payload schema.TelemetryPayload
	if err := json.Unmarshal(msg, &payload); err != nil {
		log.Printf("Worker failed to parse message: %v", err)
		return
	}

	payloadStr, _ := json.Marshal(payload)
	resultStr, err := analyzer.AnalyzeTelemetry(ctx, string(payloadStr))
	if err != nil {
		log.Printf("AI Analysis failed for event: %v", err)
		return
	}

	var result struct {
		IsZeroDay   bool    `json:"isZeroDay"`
		ThreatLevel float64 `json:"threatLevel"`
		Explanation string  `json:"explanation"`
	}

	if err := json.Unmarshal([]byte(resultStr), &result); err != nil {
		log.Printf("Failed to decode AI response analysis: %v", err)
		return
	}

	if result.IsZeroDay || result.ThreatLevel > 0.8 {
		log.Printf("[ALERT] High Threat Detected! ZeroDay: %t | ThreatLevel: %.2f", result.IsZeroDay, result.ThreatLevel)
		log.Printf("Explanation: %s", result.Explanation)
	} else {
		log.Printf("[SUPPRESSED] Fraudulent or benign alert filtered.")
	}
}
