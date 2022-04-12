package hash

import (
	"crypto/sha256"
	sha256simd "github.com/minio/sha256-simd"
	"github.com/zeebo/blake3"
	"github.com/zeebo/xxh3"
	"hash"
	"testing"
)

func benchmarkSize(b *testing.B, create func() hash.Hash, size int) {
	var bench = create()
	var buf = make([]byte, size)
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
	}
}

func BenchmarkHash(b *testing.B) {
	algos := []struct {
		name   string
		create func() hash.Hash
	}{
		{
			name:   "sha256-go",
			create: sha256.New,
		},
		{
			name: "zeebo-xxh3",
			create: func() hash.Hash {
				return xxh3.New()
			},
		},
		{
			name: "zeebo-blake3",
			create: func() hash.Hash {
				return blake3.New()
			},
		},
		{
			name: "minio-sha256simd",
			create: func() hash.Hash {
				return sha256simd.New()
			},
		},
	}

	sizes := []struct {
		n string
		f func(*testing.B, func() hash.Hash, int)
		s int
	}{
		{"8Bytes", benchmarkSize, 1 << 3},
		{"1K", benchmarkSize, 1 << 10},
		{"8K", benchmarkSize, 1 << 13},
		{"1M", benchmarkSize, 1 << 20},
		{"5M", benchmarkSize, 5 << 20},
		{"10M", benchmarkSize, 5 << 21},
	}

	for _, a := range algos {
		for _, y := range sizes {
			s := a.name + "/" + y.n
			b.Run(s, func(b *testing.B) { y.f(b, a.create, y.s) })
		}
	}
}
