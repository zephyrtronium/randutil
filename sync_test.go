package randutil

import (
	"math/rand"
	"testing"
	"time"
)

func TestSyncIdempotency(t *testing.T) {
	seed := time.Now().UnixNano()
	rnga, rngb := rand.NewSource(seed), Synchronize(rand.NewSource(seed))
	for i := 0; i < 10000; i++ {
		if a, b := rnga.Int63(), rngb.Int63(); a != b {
			t.Error("mismatch after", i, "iterations")
		}
	}
}

func TestSynchrony(t *testing.T) {
	rng := Synchronize(rand.NewSource(time.Now().UnixNano()))
	ch := make(chan int64, 1024)
	for i := 0; i < 10000; i++ {
		go func() { ch <- rng.Int63() }()
	}
	for i := 0; i < 10000; i++ {
		<-ch
	}
}