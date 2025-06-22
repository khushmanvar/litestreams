package benchmark

import (
	"os"
	"testing"
	"https://github.com/khushmanvar/litestreams/config"
)

func BenchmarkLoadConfig_Default(b *testing.B) {
	os.Unsetenv("STREAM_SHARDS")
	for i := 0; i < b.N; i++ {
		_ = config.LoadConfig()
	}
}

func BenchmarkLoadConfig_EnvOverride(b *testing.B) {
	os.Setenv("STREAM_SHARDS", "8")
	for i := 0; i < b.N; i++ {
		_ = config.LoadConfig()
	}
	os.Unsetenv("STREAM_SHARDS")
} 