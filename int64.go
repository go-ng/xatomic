package xatomic

import (
	"sync/atomic"
	"unsafe"
)

type Int64 Integer[int64]

var _ Abstract[int64] = (*Int64)(nil)

func Int64FromStd(v *atomic.Int64) *Int64 {
	return (*Int64)(unsafe.Pointer(v))
}

func Int64FromPointer(v *int64) *Int64 {
	return (*Int64)(unsafe.Pointer(v))
}

func (v *Int64) Generic() *Integer[int64] {
	return (*Integer[int64])(v)
}

func (v *Int64) Std() *atomic.Int64 {
	return (*atomic.Int64)(unsafe.Pointer(v))
}

func (v *Int64) Load() int64 {
	return atomic.LoadInt64((*int64)(&v.Value))
}

func (v *Int64) Store(new int64) {
	atomic.StoreInt64((*int64)(&v.Value), new)
}

func (v *Int64) Add(delta int64) int64 {
	return atomic.AddInt64((*int64)(&v.Value), delta)
}

func (v *Int64) Swap(new int64) int64 {
	return atomic.SwapInt64((*int64)(&v.Value), new)
}

func (v *Int64) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64((*int64)(&v.Value), old, new)
}
