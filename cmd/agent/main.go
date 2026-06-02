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

	for {
		payload := collectTelemetry()
		
		// Fix: Changed 'data' to '_' to ignore the unused variable
		_, err := json.Marshal(payload)
		if err != nil {
			log.Printf("Failed to marshal payload: %v", err)
			continue
		}

		// Logic for quarantinefile and KillProcess is commented out
		// as they are not yet defined in the codebase.
		// quarantinefile(payload) 
		// KillProcess(payload)    

		fmt.Printf("[Agent] Pushed event %s to queue\n", payload.EventID)
		time.Sleep(2 * time.Second)
	}
}

func collectTelemetry() schema.TelemetryPayload {
	return schema.TelemetryPayload{
		EventID:       fmt.Sprintf("evt-%d", time.Now().UnixNano()),
		Timestamp:     time.Now(),
		SourceNode:    "node-alpha",
		EventType:     "process_execution",
		ProcessName:   "unknown_binary.exe",
		DestinationIP: "192.168.1.100",
		PayloadData:   "Raw telemetry data",
	}
}
