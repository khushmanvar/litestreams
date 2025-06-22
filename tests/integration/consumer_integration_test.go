package integration

import (
	"strings"
	"testing"
	"time"
	"https://github.com/khushmanvar/litestreams/internal/broker"
	"https://github.com/khushmanvar/litestreams/internal/processor"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func TestConsumerIntegration(t *testing.T) {
	b := broker.NewBroker(1)
	b.Start()
	defer b.Stop()

	rec := types.Record{ID: "c-key", Payload: "hello"}
	b.RecordChan <- rec

	proc := processor.Processor{Transform: strings.ToUpper}
	var records []types.Record
	for i := 0; i < 10; i++ {
		records = b.GetRecords("c-key")
		if len(records) > 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if len(records) == 0 {
		t.Fatal("expected to get records from broker")
	}
	out := proc.Process(records[0].Payload)
	if out != "HELLO" {
		t.Errorf("expected processed payload to be HELLO, got %s", out)
	}
} 