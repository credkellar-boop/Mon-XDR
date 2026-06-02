package schema

import "time"

// TelemetryPayload represents the standardized data structure collected by your agents.
type TelemetryPayload struct {
	EventID       string    `json:"event_id"`
	Timestamp     time.Time `json:"timestamp"`
	SourceNode    string    `json:"source_node"`
	EventType     string    `json:"event_type"` // e.g., "network_flow", "process_creation"
	ProcessName   string    `json:"process_name,omitempty"`
	DestinationIP string    `json:"destination_ip,omitempty"`
	PayloadData   string    `json:"payload_data"` // Raw data or base64 encoded payload
}

// ThreatAnalysisResult represents the verdict returned by the Gemini AI analyzer.
type ThreatAnalysisResult struct {
	EventID     string  `json:"event_id"`
	ThreatLevel float64 `json:"threat_level"` // 0.0 to 1.0 (1.0 being critical)
	IsZeroDay   bool    `json:"is_zero_day"`
	Explanation string  `json:"explanation"`
}
