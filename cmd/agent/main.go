package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/credkellar-boop/Mon-XDR/pkg/action"
)

func main() {
	log.Println("Starting Mon-XDR Agent...")

	for {
		payload := collectTelemetry()

		// Use the JSON data for intelligent analysis
		data, _ := json.Marshal(payload)

		// Logic to trigger actions based on high threat score
		if payload.ProcessName == "malicious_binary.exe" {
			action.QuarantineFile(payload.PayloadData)
			action.KillProcess(payload.ProcessName)
		}

		log.Printf("[Agent] Processed event: %s", string(data))
		time.Sleep(2 * time.Second)
	}
}
