# Mon-XDR

[![Go Version](https://img.shields.io/github/go-mod/go-version/credkellar-boop/mon-xdr)](https://go.dev/)
[![Build Status](https://github.com/credkellar-boop/mon-xdr/actions/workflows/ci.yml/badge.svg)](https://github.com/credkellar-boop/mon-xdr/actions)
[![License](https://img.shields.io/github/license/credkellar-boop/mon-xdr)](LICENSE)

A Go-based Extended Detection and Response (XDR) architecture unifying high-speed endpoint telemetry with cloud infrastructure monitoring.

## Overview
Mon-XDR correlates cross-domain telemetry using Gemini's AI to neutralize zero-day threats. It features:
* **High-speed telemetry ingestion** via Go channels and workers.
* **Automated threat response** (Quarantine/KillProcess).
* **AI-driven analysis** leveraging Gemini API.
* **Kubernetes-native deployment** for scalability.

## Project Structure
```text
mon-xdr/
├── cmd/
│   ├── agent/          # Endpoint agent implementation
│   └── worker/         # Backend analysis worker
├── pkg/
│   ├── action/         # Execution logic (Quarantine/Kill)
│   ├── db/             # Persistence/Whitelisting store
│   ├── gemini/         # AI analysis integration
│   ├── ratelimit/      # API quota management
│   └── schema/         # Standardized data structures
├── deployments/        # Kubernetes manifests
├── docker-compose.yml  # Local development orchestration
└── go.mod              # Project dependencies
