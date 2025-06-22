package unit

import (
	"testing"
	"strings"
	"https://github.com/khushmanvar/litestreams/internal/processor"
	"https://github.com/khushmanvar/litestreams/internal/types"
)

func TestProcessor_Process(t *testing.T) {
	p := processor.Processor{Transform: strings.ToUpper}
	out := p.Process("foo")
	if out != "FOO" {
		t.Errorf("expected FOO, got %s", out)
	}
	p2 := processor.Processor{}
	if p2.Process("bar") != "bar" {
		t.Errorf("expected identity transform")
	}
}

func TestProcessor_ProcessWindow(t *testing.T) {
	p := processor.Processor{Transform: strings.ToLower}
	records := []types.Record{{Payload: "A"}, {Payload: "B"}}
	out := p.ProcessWindow(records)
	if !strings.Contains(out, "a;") || !strings.Contains(out, "b;") {
		t.Errorf("expected lowercased payloads in window result, got %s", out)
	}
} 