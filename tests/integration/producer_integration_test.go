package integration

import (
	"testing"
	"time"
	"https://github.com/khushmanvar/litestreams/internal/broker"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func TestProducerIntegration(t *testing.T) {
	b := broker.NewBroker(2)
	b.Start()
	defer b.Stop()

	rec := types.Record{ID: "int-key", Payload: "int-payload"}
	b.RecordChan <- rec
	var got []types.Record
	for i := 0; i < 10; i++ {
		got = b.GetRecords("int-key")
		if len(got) > 0 && got[0].Payload == "int-payload" {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if len(got) == 0 || got[0].Payload != "int-payload" {
		t.Errorf("expected to get produced record, got %+v", got)
	}
} 