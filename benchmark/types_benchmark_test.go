package benchmark

import (
	"testing"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func BenchmarkRecord_Metadata(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rec := types.Record{
			ID:        "id",
			Timestamp: 123,
			Payload:   "payload",
			Metadata:  map[string]string{"foo": "bar"},
		}
		_ = rec.Metadata["foo"]
	}
} 