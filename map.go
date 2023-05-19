package xatomic

import (
	"sync/atomic"
	"unsafe"

	"golang.org/x/exp/constraints"
)

// StoreMap atomically stores a map, given a pointer to it.
// It is a high-efficient method, because `map` is a actually
// just a pointer to a structure, so this is effectively just
// an atomic.StorePointer call.
func StoreMap[K constraints.Ordered, V any](dst *map[K]V, src map[K]V) {
	// This function works under assumption that type `map` is a pointer.
	//
	// This assumption is validated in map_test.go.
	atomic.StorePointer(
		(*unsafe.Pointer)((unsafe.Pointer)(dst)),
		(unsafe.Pointer)(unref(unsafe.Pointer(&src))),
	)
}

// LoadMap atomically loads a map, given a pointer to it.
// It is a high-efficient method, because `map` is a actually
// just a pointer to a structure, so this is effectively just
// an atomic.LoadPointer call.
func LoadMap[K constraints.Ordered, V any](src *map[K]V) map[K]V {
	// This function works under assumption that type `map` is a pointer.
	//
	// This assumption is validated in map_test.go.
	return *(*map[K]V)(ref(atomic.LoadPointer(
		(*unsafe.Pointer)((unsafe.Pointer)(src))),
	))
}
