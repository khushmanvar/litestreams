package unit

import (
	"testing"
	"time"
	"https://github.com/khushmanvar/litestreams/internal/broker"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func TestBroker_AppendAndGetRecords(t *testing.T) {
	b := broker.NewBroker(2)
	b.Start()
	defer b.Stop()

	rec1 := types.Record{ID: "foo", Payload: "bar"}
	rec2 := types.Record{ID: "baz", Payload: "qux"}
	b.AppendRecord(rec1)
	b.AppendRecord(rec2)

	got := b.GetRecords("foo")
	if len(got) == 0 || got[0].ID != "foo" {
		t.Errorf("expected to get record with ID 'foo', got %+v", got)
	}

	got2 := b.GetRecords("baz")
	if len(got2) == 0 || got2[0].ID != "baz" {
		t.Errorf("expected to get record with ID 'baz', got %+v", got2)
	}
}

func TestBroker_Retention(t *testing.T) {
	b := broker.NewBrokerWithRetention(1, 3)
	b.Start()
	defer b.Stop()
	for i := 0; i < 10; i++ {
		rec := types.Record{ID: "foo", Payload: string(rune('a' + i))}
		b.AppendRecord(rec)
	}
	got := b.GetRecords("foo")
	if len(got) != 3 {
		t.Errorf("expected 3 records due to retention, got %d", len(got))
	}
	if got[0].Payload != "h" || got[2].Payload != "j" {
		t.Errorf("unexpected records after retention: %+v", got)
	}
}

func TestBroker_ChannelIngestion(t *testing.T) {
	b := broker.NewBroker(1)
	b.Start()
	defer b.Stop()
	rec := types.Record{ID: "foo", Payload: "bar"}
	b.RecordChan <- rec
	var got []types.Record
	for i := 0; i < 10; i++ {
		got = b.GetRecords("foo")
		if len(got) > 0 && got[0].Payload == "bar" {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}
	if len(got) == 0 || got[0].Payload != "bar" {
		t.Errorf("expected to get record via channel, got %+v", got)
	}
} 