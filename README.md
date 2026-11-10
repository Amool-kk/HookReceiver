# HookReceiver

## Project objective
- HookReceiver is a lightweight webhook inspection service written in Go. It aims to provide developers with a simple way to receive, inspect and debug webhook events.

## Current Implementation Status
The project is in an early bootstrap stage. Currently implemented:
- A basic HTTP server built with the Go standard library (`net/http`), no external dependencies.
- Routing via `http.ServeMux` using Go 1.22+ method-based patterns.
- A single endpoint:

| Method | Path | Description |
|--------|------|-------------|
| `GET`  | `/`  | Returns a welcome message: `Welcome to Hook Receiver` |

- Server address is provided by a `config` struct via `mustLoadConfig()`; currently defaults to `127.0.0.1:8080` (hardcoded).
- Structured logging using `log/slog`.
- Graceful shutdown on `SIGINT` / `SIGTERM` (e.g. `Ctrl+C`), allowing in-flight requests up to 5 seconds to complete.

## Prerequisites
- [Go](https://go.dev/dl/) 1.26.3 or later (as specified in `go.mod`)
- Git

## Installation Instructions
```bash
git clone https://github.com/Amool-kk/hookreceiver.git
cd hookreceiver
go mod download
```

## How to Run the Application
Run directly:
```bash
go run ./cmd/hookreceiver
```

Or build a binary and run it:
```bash
go build -o bin/hookreceiver ./cmd/hookreceiver
./bin/hookreceiver
```

Verify it is running:
```bash
curl http://127.0.0.1:8080/
# Welcome to Hook Receiver
```

Stop the server with `Ctrl+C`; it will shut down gracefully.

## Project Structure
```
.
├── cmd/
│   └── hookreceiver/
│       └── main.go      # Entry point: config, HTTP server, graceful shutdown
├── docs/
│   ├── architecture/    # Architecture documentation (placeholder)
│   └── quality/         # Quality documentation (placeholder)
├── go.mod
└── README.md
```

## Current Limitations
- No webhook receiving endpoint yet — only a welcome route on `GET /`.
- No storage or inspection of incoming requests (headers, body, query params).
- No UI or API for viewing captured webhook events.
- `mustLoadConfig()` returns a hardcoded address (`127.0.0.1:8080`); no configuration via flags or environment variables yet.
- Binds to localhost only, so it is not reachable from external webhook providers without a tunnel.
- No HTTP server timeouts (`ReadTimeout`, `WriteTimeout`, `IdleTimeout`) configured.
- No tests.
- `docs/` directories are empty placeholders.