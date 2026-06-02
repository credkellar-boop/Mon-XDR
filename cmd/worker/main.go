package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/credkellar-boop/Mon-XDR/kg/gemini"
	"github.com/credkellar-boop/Mon-XDR/pkg/schema"
)

func main() {
	log.Println("Starting Mon-XDR Worker...")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Initialize the Gemini Analyzer
	analyzer, err := gemini.NewAnalyzer(ctx)
	if err != nil {
		log.Fatalf("Failed to initialize Gemini analyzer: %v", err)
	}
	defer analyzer.Close()

	// TODO: Initialize your queue consumer here
	// e.g., consumer := InitQueueConsumer("telemetry_topic")
	// msgs := consumer.Consume()

	// Simulated incoming message channel for demonstration
	msgs := make(chan []byte)

	// Graceful shutdown handling
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

func processMessage(ctx context.Context, analyzer *gemini.Analyzer, msg []byte) {
	var payload schema.TelemetryPayload
	if err := json.Unmarshal(msg, &payload); err != nil {
		log.Printf("Worker failed to parse message: %v", err)
		return
	}

	// Send to Gemini for cross-domain analysis
	result, err := analyzer.AnalyzeTelemetry(ctx, payload)
	if err != nil {
		log.Printf("AI Analysis failed for event %s: %v", payload.EventID, err)
		return
	}

	// Logic to aggressively filter and neutralize threats or fraudulent alerts
	if result.IsZeroDay || result.ThreatLevel > 0.8 {
		log.Printf("[ALERT] High Threat Detected! EventID: %s | ZeroDay: %v | ThreatLevel: %.2f", 
			result.EventID, result.IsZeroDay, result.ThreatLevel)
		log.Printf("Explanation: %s", result.Explanation)
		// TODO: Trigger orchestration response (e.g., isolate node, kill process)
	} else {
		// Suppress fraudulent/benign alerts to keep the pipeline clean
		log.Printf("[SUPPRESSED] Fraudulent or benign alert blocked. EventID: %s", result.EventID)
	}
}
