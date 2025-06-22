package types

// Purpose: Define shared data structures used across components

type Record struct {
    ID        string
    Timestamp int64
    Payload   string
    Metadata  map[string]string // Additional metadata fields
}