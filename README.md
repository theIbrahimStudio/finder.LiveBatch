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

Via command-line flags:

```bash
go run main.go --max-batch-size=4 --max-latency-ms=100 --listen-addr=":9000"
```

Via environment variables:

```bash
LIVEBATCH_MAX_BATCH_SIZE=16 LIVEBATCH_MAX_LATENCY_MS=200 go run main.go
```

---

## Configuration

LiveBatch supports configuration via **environment variables** or **command-line flags**, using `viper` + `pflag`.

| Name              | Flag               | Env Var                    | Default | Description                      |
| ----------------- | ------------------ | -------------------------- | ------- | -------------------------------- |
| Max Batch Size    | `--max-batch-size` | `LIVEBATCH_MAX_BATCH_SIZE` | `8`     | Max number of requests per batch |
| Max Latency (ms)  | `--max-latency-ms` | `LIVEBATCH_MAX_LATENCY_MS` | `50`    | Max wait time before dispatching |
| Listening Address | `--listen-addr`    | `LIVEBATCH_LISTEN_ADDR`    | `:8080` | HTTP server bind address         |

---

## Roadmap

- [x] HTTP dynamic batching proxy (MVP)
- [x] Config via environment or CLI
- [ ] gRPC and ONNX backend support
- [ ] Prometheus metrics
- [ ] Deadline-based and priority queueing
- [ ] Docker + Helm chart for Kubernetes
- [ ] Python client SDK

---

## Contributing

PRs welcome! Check out the [CONTRIBUTING.md](https://github.com/theIbrahimStudio/.github/blob/main/CONTRIBUTING.md) for guidelines.
