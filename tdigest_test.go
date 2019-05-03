package tdigest

import (
	"math/rand"
	"strconv"
	"testing"
)

func TestTDigest(t *testing.T) {
	td := NewConcurrent(Compression(123))
	if td.ValueAt(1) != 0 {
		t.Fatalf("failed to retreive value from empty TDigest, %v", td.ValueAt(1))
	}
	td.Add(1, 1)
	td.Add(2, 1)
	if val := td.ValueAt(0); val != 1 {
		t.Fatalf("Unexpected quantile of large value: got %v, expected %v",
			val, 1)
	}
}

func BenchmarkAdd(b *testing.B) {
	for _, size := range []int{10, 100, 200, 500, 1000, 5000, 10000, 50000} {
		newSketch := func() Sketch { return New(Compression(float64(size))) }
		b.Run(strconv.Itoa(size), benchmarkAddSize(newSketch, size))
	}
}

func BenchmarkConcurrentAdd(b *testing.B) {
	for _, size := range []int{10, 100, 200, 500, 1000, 5000, 10000, 50000} {
		newSketch := func() Sketch {
			return NewConcurrent(Compression(float64(size)))
		}
		b.Run(strconv.Itoa(size), benchmarkAddSize(newSketch, size))
	}
}

func benchmarkAddSize(newSketch func() Sketch, size int) func(*testing.B) {
	return func(b *testing.B) {
		h := newSketch()
		data := make([]float64, 0, b.N)
		for i := 0; i < b.N; i++ {
			data = append(data, rand.Float64())
		}
		b.ResetTimer()
		for _, v := range data {
			h.Add(v, 1)
		}
	}
}
