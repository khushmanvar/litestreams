package main

import (
    "fmt"
    "time"
    "https://github.com/khushmanvar/litestreams/config"
    "https://github.com/khushmanvar/litestreams/internal/broker"
    "https://github.com/khushmanvar/litestreams/internal/types"
)

// Purpose: This file simulates a data producer that pushes data to the stream.

func main() {
    cfg := config.LoadConfig()
    b := broker.NewBroker(cfg.StreamShards)
    b.Start()

    for i := 0; i < 5; i++ {
        rec := types.Record{
            ID:        fmt.Sprintf("key-%d", i),
            Timestamp: time.Now().UnixNano(),
            Payload:   fmt.Sprintf("payload-%d", i),
        }
        b.RecordChan <- rec
        fmt.Printf("Produced: %+v\n", rec)
        time.Sleep(200 * time.Millisecond)
    }

    // Give broker time to ingest
    time.Sleep(1 * time.Second)
    b.Stop()
}
