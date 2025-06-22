package unit

import (
	"os"
	"testing"
	"https://github.com/khushmanvar/litestreams/config"
)

func TestLoadConfig_Default(t *testing.T) {
	os.Unsetenv("STREAM_SHARDS")
	cfg := config.LoadConfig()
	if cfg.StreamShards != 4 {
		t.Errorf("expected default StreamShards=4, got %d", cfg.StreamShards)
	}
}

func TestLoadConfig_EnvOverride(t *testing.T) {
	os.Setenv("STREAM_SHARDS", "7")
	cfg := config.LoadConfig()
	if cfg.StreamShards != 7 {
		t.Errorf("expected StreamShards=7 from env, got %d", cfg.StreamShards)
	}
	os.Unsetenv("STREAM_SHARDS")
} 