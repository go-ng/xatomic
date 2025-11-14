package xatomic

// Value is a wrapper around a customly typed value that provides atomic operations.
type Value[T any] struct {
	Pointer[T]
}

// Load returns the value set by the most recent Store.
// It returns zero value if there has been no call to Store for this Value.
func (v *Value[T]) Load() T {
	ptr := v.Pointer.Load()
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

// Store sets the value of the [Value] v to val.
func (v *Value[T]) Store(new T) {
	v.Pointer.Store(&new)
}

// Swap stores new into Value and returns the previous value. It returns zero value
// if the Value is empty.
func (v *Value[T]) Swap(new T) (old T) {
	curPtr := v.Pointer.Swap(&new)
	if curPtr == nil {
		var zero T
		return zero
	}
	return *curPtr
}
