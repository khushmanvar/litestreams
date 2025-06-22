package broker

import (
    "sync"
    "hash/fnv"
    "https://github.com/khushmanvar/litestreams/internal/types"
)

const defaultRetention = 100 // max records per shard

// Purpose: In-memory broker simulating stream partitions and record retention

// shard represents a single partition's append-only log.
type shard struct {
    mu      sync.RWMutex
    records []types.Record
}

type Broker struct {
    shards     []*shard
    retention  int
    RecordChan chan types.Record // Channel for async ingestion
    stopCh     chan struct{}
}

// NewBroker creates a broker with the given number of shards.
func NewBroker(numShards int) *Broker {
    return NewBrokerWithRetention(numShards, defaultRetention)
}

// NewBrokerWithRetention allows custom retention per shard.
func NewBrokerWithRetention(numShards, retention int) *Broker {
    shards := make([]*shard, numShards)
    for i := range shards {
        shards[i] = &shard{}
    }
    return &Broker{
        shards:     shards,
        retention:  retention,
        RecordChan: make(chan types.Record, 100),
        stopCh:     make(chan struct{}),
    }
}

// Start launches a goroutine to ingest records from RecordChan.
func (b *Broker) Start() {
    go func() {
        for {
            select {
            case rec := <-b.RecordChan:
                b.AppendRecord(rec)
            case <-b.stopCh:
                return
            }
        }
    }()
}

// Stop signals the broker to stop ingesting from the channel.
func (b *Broker) Stop() {
    close(b.stopCh)
}

// getShard returns the shard for a given key.
func (b *Broker) getShard(key string) *shard {
    h := fnv.New32a()
    h.Write([]byte(key))
    idx := int(h.Sum32()) % len(b.shards)
    return b.shards[idx]
}

// AppendRecord appends a record to the appropriate shard based on key, enforcing retention.
func (b *Broker) AppendRecord(record types.Record) {
    s := b.getShard(record.ID)
    s.mu.Lock()
    defer s.mu.Unlock()
    s.records = append(s.records, record)
    if len(s.records) > b.retention {
        // Remove oldest records to enforce retention
        s.records = s.records[len(s.records)-b.retention:]
    }
}

// GetRecords returns all records from the shard for a given key.
func (b *Broker) GetRecords(key string) []types.Record {
    s := b.getShard(key)
    s.mu.RLock()
    defer s.mu.RUnlock()
    // Return a copy to avoid race conditions
    out := make([]types.Record, len(s.records))
    copy(out, s.records)
    return out
}