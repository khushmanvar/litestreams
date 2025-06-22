package benchmark

import (
	"strings"
	"testing"
	"https://github.com/khushmanvar/litestreams/internal/processor"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func BenchmarkProcessor_Process(b *testing.B) {
	p := processor.Processor{Transform: strings.ToUpper}
	for i := 0; i < b.N; i++ {
		_ = p.Process("payload")
	}
}

func BenchmarkProcessor_ProcessWindow(b *testing.B) {
	p := processor.Processor{Transform: strings.ToLower}
	records := make([]types.Record, 1000)
	for i := range records {
		records[i] = types.Record{Payload: "PAYLOAD"}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = p.ProcessWindow(records)
	}
} 