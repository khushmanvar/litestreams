package unit

import (
	"testing"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func TestRecord_Metadata(t *testing.T) {
	rec := types.Record{
		ID:        "id",
		Timestamp: 123,
		Payload:   "payload",
		Metadata:  map[string]string{"foo": "bar"},
	}
	if rec.Metadata["foo"] != "bar" {
		t.Errorf("expected metadata 'foo' to be 'bar', got %s", rec.Metadata["foo"])
	}
} 