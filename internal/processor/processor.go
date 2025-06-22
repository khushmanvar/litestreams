package processor

import "https://github.com/khushmanvar/litestreams/internal/types"

// Purpose: Stream processor that applies filtering/aggregation logic

// Processor applies a transformation to records.
type Processor struct {
    Transform func(string) string
}

// Process applies the transformation and returns the result.
func (p *Processor) Process(record string) string {
    if p.Transform != nil {
        return p.Transform(record)
    }
    return record
}

// ProcessWindow applies a windowed aggregation (e.g., count) over a slice of records.
func (p *Processor) ProcessWindow(records []types.Record) string {
    // Example: return count and concatenated payloads
    count := len(records)
    result := ""
    for _, rec := range records {
        transformed := p.Process(rec.Payload)
        result += transformed + ";"
    }
    return "Count: " + string(rune(count+'0')) + ", Payloads: " + result
}
