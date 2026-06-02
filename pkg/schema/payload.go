type TelemetryPayload struct {
    AgentID   string `json:"agent_id"`
    Timestamp string `json:"timestamp"`
    Source    string `json:"source"`    // "ENDPOINT" or "CLOUD"
    EventType string `json:"event_type"`
    Data      string `json:"data"`       // The raw hash or log text
}
