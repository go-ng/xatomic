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

// CompareAndSwapPointer atomically compares and swaps a pointer to a value.
// This is just a type-safe wrapper for atomic.CompareAndSwapPointer.
//
// Consider also using atomic.Pointer, instead of this function.
func CompareAndSwapPointer[T any](addr **T, old, new *T) bool {
	return atomic.CompareAndSwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(addr)),
		unsafe.Pointer(old),
		unsafe.Pointer(new),
	)
}

// CompareAndSwapPointer atomically compares and swaps a pointer to a value.
// This is just a type-safe wrapper for atomic.SwapPointer.
//
// Consider also using atomic.Pointer, instead of this function.
func SwapPointer[T any](addr **T, new *T) *T {
	return (*T)(atomic.SwapPointer(
		(*unsafe.Pointer)(unsafe.Pointer(addr)),
		unsafe.Pointer(new),
	))
}
