package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/credkellar-boop/Mon-XDR/pkg/schema"
)

func main() {
	log.Println("Starting Mon-XDR Agent...")

	// TODO: Initialize your connection to the asynchronous queue here.
	// e.g., queueClient := InitQueue()
	// defer queueClient.Close()

	// Simulated telemetry collection loop
	for {
		payload := collectTelemetry()
		
		// Serialize to JSON for the queue
		data, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Failed to marshal payload: %v", err)
			continue
		}

		// TODO: Push 'data' to the queue
		// queueClient.Publish("telemetry_topic", data)
		fmt.Printf("[Agent] Pushed event %s to queue\n", payload.EventID)

		time.Sleep(2 * time.Second) // Poll interval
	}
}

// collectTelemetry mocks the gathering of system data.
// In production, this would hook into eBPF, Windows Event Logs, or network interfaces.
func collectTelemetry() schema.TelemetryPayload {
	return schema.TelemetryPayload{
		EventID:       fmt.Sprintf("evt-%d", time.Now().UnixNano()),
		Timestamp:     time.Now(),
		SourceNode:    "node-alpha",
		EventType:     "process_execution",
		ProcessName:   "unknown_binary.exe",
		DestinationIP: "192.168.1.100",
		PayloadData:   "Raw command line arguments or hash data here",
	}
}
