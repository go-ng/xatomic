package xatomic

import (
	"sync/atomic"
)

// Value is a wrapper around a customly typed value that provides atomic operations.
type Value[T any] struct {
	Value atomic.Value
}

// Load returns the value set by the most recent Store.
// It returns zero value if there has been no call to Store for this Value.
func (v *Value[T]) Load() T {
	cur := v.Value.Load()
	if cur == nil {
		var zero T
		return zero
	}
	return cur.(T)
}

// Store sets the value of the [Value] v to val.
func (v *Value[T]) Store(new T) {
	v.Value.Store(new)
}

// Swap stores new into Value and returns the previous value. It returns zero value
// if the Value is empty.
func (v *Value[T]) Swap(new T) (old T) {
	cur := v.Value.Swap(new)
	if cur == nil {
		var zero T
		return zero
	}
	return cur.(T)
}

// CompareAndSwap executes the compare-and-swap operation for the [Value].
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.Value.CompareAndSwap(old, new)
}
