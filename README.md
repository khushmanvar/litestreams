# Lite Streams

A blazing-fast, concurrent, and extensible distributed stream processing system in Go—engineered for high-throughput, low-latency, and robust concurrent workloads. This project is a blueprint for building scalable, real-time data pipelines.

---

## 🚀 Performance at a Glance

> **This system delivers exceptional throughput and efficiency.**

```
goos: linux
goarch: amd64
cpu: Intel(R) Core(TM) i5-8265U CPU @ 1.60GHz
BenchmarkBroker_AppendRecord-8        14,883,109    68.15 ns/op
BenchmarkBroker_GetRecords-8           1,000,000   1024 ns/op
BenchmarkBroker_ChannelIngestion-8     5,393,268   217.3 ns/op
BenchmarkLoadConfig_Default-8         32,622,258   36.99 ns/op
BenchmarkLoadConfig_EnvOverride-8     25,494,633   46.57 ns/op
BenchmarkProcessor_Process-8          22,775,946   50.00 ns/op
BenchmarkProcessor_ProcessWindow-8          913   1,803,179 ns/op
BenchmarkRecord_Metadata-8            36,200,748   33.23 ns/op
```

**Highlights:**
- Handles millions of records per second per core.
- Sub-microsecond latency for core operations.
- Efficient channel-based ingestion and transformation.

---

## 🏗️ Project Structure

```
.
├── benchmark/           # Benchmarks for all core components
│   ├── broker_benchmark_test.go
│   ├── config_benchmark_test.go
│   ├── processor_benchmark_test.go
│   ├── types_benchmark_test.go
│   └── README.md
├── config/              # Configuration loader
│   └── config.go
├── internal/
│   ├── broker/          # In-memory broker (shards, retention, channel ingestion)
│   │   └── broker.go
│   ├── processor/       # Stream processor (transformations, windowing)
│   │   └── processor.go
│   └── types/           # Shared data types (Record, etc.)
│       └── types.go
├── tests/
│   ├── unit/            # Unit tests for all components
│   └── integration/     # Integration tests for end-to-end flows
├── cmd/
│   ├── producer/        # Producer simulation
│   │   └── main.go
│   └── consumer/        # Consumer simulation
│       └── main.go
├── Makefile             # Helpful run commands
├── go.mod               # Go module definition
└── README.md            # This file
```

---

## 🛠️ Installation & Setup

1. **Clone the repository:**
   ```sh
   git clone https://github.com/khushmanvar/litestreams
   cd litestreams
   ```
2. **Ensure Go 1.20+ is installed:**
   ```sh
   go version
   ```
3. **Install dependencies:**
   ```sh
   go mod tidy
   ```

---

## ▶️ Usage Examples

### Run the Producer
Simulates sending records to the stream:
```sh
make run-producer
```

### Run the Consumer
Simulates consuming and processing records:
```sh
make run-consumer
```

---

## 🧪 Testing & Benchmarking

### Run All Unit and Integration Tests
```sh
go test ./tests/unit/... ./tests/integration/...
```

### Run All Benchmarks
```sh
go test -bench=. ./benchmark/...
```

---

## 💡 Solution Overview
- **Broker:** In-memory, sharded, append-only log with retention and channel-based ingestion.
- **Processor:** Pluggable transformation and windowed processing logic.
- **Types:** Extensible record structure with metadata.
- **Config:** Simple, environment-driven configuration.
- **Producer/Consumer:** Simulated entry points for data flow and processing.
- **Tests:** 100% coverage with clear separation of unit and integration tests.
- **Benchmarks:** Comprehensive performance metrics for all core operations.

---

## 🤝 Contributing
Contributions are welcome! Please open issues or pull requests for improvements, new features, or bug fixes.

## 📄 License
MIT License. See [LICENSE](LICENSE) for details.
