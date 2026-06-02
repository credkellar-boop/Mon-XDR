# Mon-XDR
a Go-based Extended Detection and Response (XDR) architecture unifying high-speed endpoint telemetry with cloud infrastructure monitoring. Utilizing secure API gateways, asynchronous task queues, and Docker orchestration, it leverages Gemini's AI to correlate cross-domain telemetry and neutralize zero-day threats across all environments.

mon-edr/
├── cmd/
│   ├── agent/             # Local endpoint binary
│   │   └── main.go        # Handles hashing, file watching, and local quarantine
│   ├── cloud-aws/         # AWS log poller
│   │   └── main.go        # Polls CloudTrail/IAM logs
│   └── worker/            # The Gemini AI analyzer
│       └── main.go        # Pulls from queue + rate-limited Gemini calls
├── internal/
│   ├── auth/              # JWT verification logic
│   ├── quarantine/        # AES encryption/isolation logic
│   └── queue/             # Redis/RabbitMQ client configuration
├── pkg/
│   ├── gemini/            # Gemini API client wrapper
│   └── schema/            # Unified struct definitions for telemetry
├── deployments/
│   └── docker-compose.yml # Orchestrates Redis, Workers, and Gateways
├── go.mod                 # Project dependencies
└── README.md
