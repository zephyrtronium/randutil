package randutil

import (
	"math/rand"
	"testing"
	"time"
)

func TestPregeneratorIdempotency(t *testing.T) {
	seed := time.Now().UnixNano()
	rnga, rngb := rand.NewSource(seed), Pregenerate(rand.NewSource(seed), 1024)
	for i := 0; i < 10000; i++ {
		if a, b := rnga.Int63(), rngb.Int63(); a != b {
			t.Error("mismatch after", i, "iterations")
		}
	}
}

func BenchmarkPregenerator(b *testing.B) {
	b.StopTimer()
	rng := Pregenerate(rand.NewSource(time.Now().UnixNano()), 1024)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		rng.Int63()
	}
}
