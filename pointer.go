package xatomic

import (
	"sync/atomic"
	"unsafe"
)

// StorePointer atomically stores a pointer to a value.
// This is just a type-safe wrapper for atomic.LoadPointer.
//
// Consider also using atomic.Pointer, instead of this function.
func StorePointer[T any](dst **T, src *T) {
	atomic.StorePointer(
		(*unsafe.Pointer)(unsafe.Pointer(dst)),
		unsafe.Pointer(src),
	)
}

// LoadPointer atomically loads a pointer to a value.
// This is just a type-safe wrapper for atomic.LoadPointer.
//
// Consider also using atomic.Pointer, instead of this function.
func LoadPointer[T any](src **T) *T {
	return (*T)(atomic.LoadPointer(
		(*unsafe.Pointer)(unsafe.Pointer(src)),
	))
}
