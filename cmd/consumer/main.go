package main

import (
    "fmt"
    "strings"
    "time"
    "https://github.com/khushmanvar/litestreams/config"
    "https://github.com/khushmanvar/litestreams/internal/broker"
    "https://github.com/khushmanvar/litestreams/internal/processor"
    "https://github.com/khushmanvar/litestreams/internal/types"
)

// Purpose: This file simulates a consumer that processes data from the stream.

func main() {
    cfg := config.LoadConfig()
    b := broker.NewBroker(cfg.StreamShards)
    b.Start()

    // Simulate records being sent asynchronously
    go func() {
        for i := 0; i < 5; i++ {
            rec := types.Record{
                ID:      fmt.Sprintf("key-%d", i),
                Payload: fmt.Sprintf("payload-%d", i),
            }
            b.RecordChan <- rec
            time.Sleep(200 * time.Millisecond)
        }
    }()

    proc := processor.Processor{
        Transform: strings.ToUpper,
    }

    // Poll for new records for a specific key
    for i := 0; i < 5; i++ {
        records := b.GetRecords("key-2")
        if len(records) > 0 {
            windowResult := proc.ProcessWindow(records)
            fmt.Printf("Windowed Processing Result for key-2: %s\n", windowResult)
        } else {
            fmt.Println("No records found for key-2")
        }
        time.Sleep(300 * time.Millisecond)
    }

    b.Stop()
}
