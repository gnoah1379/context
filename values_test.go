package context

import (
	"context"
	"testing"
)

func TestWithValues(t *testing.T) {
	parentCtx := context.WithValue(context.Background(), "k1", "v1")
	ctx := WithValues(parentCtx, KV{"k2", "v2"}, KV{"k3", 3})
	if ctx.String("k1") != "v1" {
		t.Fatal("k1 != v1")
	}
	if ctx.String("k2") != "v2" {
		t.Fatal("k2 != v2")
	}
	if ctx.String("k3") != "" {
		t.Fatal("string k3 not empty")
	}
	if ctx.Int("k3") != 3 {
		t.Fatal("k3 != 3")
	}
}
