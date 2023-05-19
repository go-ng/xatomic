package xatomic

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

type hmap struct{}

//go:linkname hmap_incrnoverflow runtime.(*hmap).incrnoverflow
func hmap_incrnoverflow(h *hmap)

func TestMapIsPointer(t *testing.T) {
	m := map[int]int{
		1: 1,
		2: 2,
	}

	t.Run("Store&Load", func(t *testing.T) {
		var storage map[int]int
		StoreMap(&storage, m)
		require.Equal(t, m, storage)
		v := LoadMap(&storage)
		require.Equal(t, m, v)
	})

	// The test above should be good enough, but just in case some additional tests
	// go below:

	t.Run("call_internal_method", func(t *testing.T) {
		// Rechecking the assumption that `map` is actually a pointer
		// (to the map header structure).
		//
		// Assuming that it may segfault or something, if something is wrong:
		hdr := (*hmap)(unref((unsafe.Pointer)(&m)))
		hmap_incrnoverflow(hdr)
	})
}

var v map[int]int
var r map[int]int

func Benchmark(b *testing.B) {
	m := map[int]int{
		1: 1,
		2: 2,
	}

	b.Run("noop", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
		}
	})

	b.Run("Map", func(b *testing.B) {
		b.Run("Store", func(b *testing.B) {
			b.Run("atomic", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					StoreMap(&v, m)
				}
			})
			b.Run("unatomic", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					v = m
				}
			})
		})
		b.Run("Load", func(b *testing.B) {
			b.Run("atomic", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					r = LoadMap(&v)
				}
			})
			b.Run("unatomic", func(b *testing.B) {
				for i := 0; i < b.N; i++ {
					r = v
				}
			})
		})
	})
}
