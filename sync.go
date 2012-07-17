package randutil

import (
	"math/rand"
	"sync"
)

type syncSource struct {
	r rand.Source
	m sync.Locker
}

// Synchronize a randomness source using a new mutex.
func Synchronize(source rand.Source) rand.Source {
	return syncSource{source, new(sync.Mutex)}
}

// Synchronize a randomness source using an existing mutex.
func SynchronizeWith(source rand.Source, mutex sync.Locker) rand.Source {
	return syncSource{source, mutex}
}

func (s syncSource) Int63() int64 {
	s.m.Lock()
	defer s.m.Unlock()
	return s.r.Int63()
}

func (s syncSource) Seed(seed int64) {
	s.m.Lock()
	defer s.m.Unlock()
	s.r.Seed(seed)
}