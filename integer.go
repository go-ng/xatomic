package xatomic

import (
	"golang.org/x/exp/constraints"
)

type AbstractInteger[T constraints.Integer] interface {
	Abstract[T]
	Add(delta T) T
}

type Integer[T constraints.Integer] struct {
	_     noCopy
	Value T
}

var _ Abstract[int32] = (*Integer[int32])(nil)

func (v *Integer[T]) typeCast() AbstractInteger[T] {
	switch v := any(v).(type) {
	case *Integer[int32]:
		return any((*Int32)(v)).(AbstractInteger[T])
	case *Integer[int64]:
		return any((*Int64)(v)).(AbstractInteger[T])
	case *Integer[uint32]:
		return any((*Uint32)(v)).(AbstractInteger[T])
	case *Integer[uint64]:
		return any((*Uint64)(v)).(AbstractInteger[T])
	default:
		panic("unsupported type")
	}
}

func (v *Integer[T]) Load() T {
	return v.typeCast().Load()
}

func (v *Integer[T]) Store(new T) {
	v.typeCast().Store(new)
}

func (v *Integer[T]) Add(delta T) T {
	return v.typeCast().Add(delta)
}

func (v *Integer[T]) Swap(new T) T {
	return v.typeCast().Swap(new)
}

func (v *Integer[T]) CompareAndSwap(old, new T) bool {
	return v.typeCast().CompareAndSwap(old, new)
}
