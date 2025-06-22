package benchmark

import (
	"testing"
	"https://github.com/khushmanvar/litestreams/internal/broker"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func BenchmarkBroker_AppendRecord(b *testing.B) {
	br := broker.NewBroker(8)
	br.Start()
	defer br.Stop()
	rec := types.Record{ID: "bench-key", Payload: "payload"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		br.AppendRecord(rec)
	}
}

func BenchmarkBroker_GetRecords(b *testing.B) {
	br := broker.NewBroker(8)
	br.Start()
	defer br.Stop()
	rec := types.Record{ID: "bench-key", Payload: "payload"}
	for i := 0; i < 1000; i++ {
		br.AppendRecord(rec)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = br.GetRecords("bench-key")
	}
}

func BenchmarkBroker_ChannelIngestion(b *testing.B) {
	br := broker.NewBroker(8)
	br.Start()
	defer br.Stop()
	rec := types.Record{ID: "bench-key", Payload: "payload"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		br.RecordChan <- rec
	}
} 