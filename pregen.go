package randutil

import "math/rand"

type pregenerator struct {
	iters  []int64
	tap    int
	source rand.Source
}

// Pregenerate iterations so that new random values are read from an array
// instead of calculated.
func Pregenerate(source rand.Source, count int) rand.Source {
	return pregenerator{make([]int64, count), count, source}
}

func (p *pregenerator) Int64() int64 {
	if p.tap < len(p.iters) {
		p.tap++
		return p.iters[p.tap-1]
	}
	for i := range p.iters {
		p.iters[i] = p.source.Int64()
	}
	p.tap = 0
	return p.iters[0]
}

func (p *pregenerator) Seed(seed int64) {
	p.source.Seed(seed)
}
