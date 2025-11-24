package xatomic

import (
	"sync/atomic"
	"unsafe"
)

type Int32 Integer[int32]

var _ Abstract[int32] = (*Int32)(nil)

func Int32FromStd(v *atomic.Int32) *Int32 {
	return (*Int32)(unsafe.Pointer(v))
}

func Int32FromPointer(v *int32) *Int32 {
	return (*Int32)(unsafe.Pointer(v))
}

func (v *Int32) Load() int32 {
	return atomic.LoadInt32((*int32)(&v.Value))
}

func (v *Int32) Store(new int32) {
	atomic.StoreInt32((*int32)(&v.Value), new)
}

func (v *Int32) Add(delta int32) int32 {
	return atomic.AddInt32((*int32)(&v.Value), delta)
}

func (v *Int32) Swap(new int32) int32 {
	return atomic.SwapInt32((*int32)(&v.Value), new)
}

func (v *Int32) CompareAndSwap(old, new int32) bool {
	return atomic.CompareAndSwapInt32((*int32)(&v.Value), old, new)
}
