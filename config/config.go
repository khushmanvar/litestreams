package config

import (
    "os"
    "strconv"
)

// Purpose: Load and store configuration values for the system

// Config holds system-wide configuration.
type Config struct {
    StreamShards int
}

// LoadConfig loads configuration from environment variables or uses defaults.
func LoadConfig() Config {
    shards := 4 // default
    if val := os.Getenv("STREAM_SHARDS"); val != "" {
        if n, err := strconv.Atoi(val); err == nil && n > 0 {
            shards = n
        }
    }
    return Config{StreamShards: shards}
}
