# LiveBatch ⚡

**LiveBatch** is a lightweight, framework-agnostic middleware that dynamically batches inference requests in real time to maximize GPU/TPU utilization. It's designed for ML teams and microservices that want to increase throughput without modifying their model code.

---

## Features

- Plug-and-play: Drop-in proxy over any HTTP ML server
- Dynamic batching: Configurable by max latency or batch size
- Written in Go: Fast, concurrent, production-ready
- Transparent: Accepts and returns single inference calls
- Efficient: Great for reducing GPU/TPU underutilization

---

## Architecture

LiveBatch acts as a sidecar or proxy in front of your model service.

```
\[Client] ---> \[LiveBatch] ---> \[Model Server]
\|  |
\|  |---> Queues requests
\|------> Batches & dispatches
```

---

## Quick Start

### Run it locally

```bash
go run main.go
```

### Test it

```bash
curl -X POST -H "Content-Type: application/json" \\
     -d '{"input": "hello"}' http://localhost:8080/infer
```

You’ll get:

```json
{ "result": "ok" }
```

---

## Configuration

| Env / Config   | Default | Description                  |
| -------------- | ------- | ---------------------------- |
| `MaxBatchSize` | `8`     | Max number of requests/batch |
| `MaxLatencyMs` | `50`    | Max delay before flushing    |
| `ListenAddr`   | `:8080` | Address to bind to           |

Config is hardcoded for now — env/config support coming soon.

---

## Roadmap

- [x] HTTP dynamic batching proxy (MVP)
- [ ] Config via environment or CLI
- [ ] gRPC and ONNX backend support
- [ ] Prometheus metrics
- [ ] Deadline-based and priority queueing
- [ ] Docker + Helm chart for Kubernetes
- [ ] Python client SDK

---

## Contributing

PRs welcome! Check out the [CONTRIBUTING.md](https://github.com/theIbrahimStudio/.github/blob/main/CONTRIBUTING.md) for guidelines.
