package xatomic

type Abstract[T any] interface {
	Load() T
	Store(new T)
	Swap(new T) T
	CompareAndSwap(old, new T) bool
}

type noCopy struct{}
