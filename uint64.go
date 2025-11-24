package xatomic

import (
	"sync/atomic"
	"unsafe"
)

type Uint64 Integer[uint64]

var _ Abstract[uint64] = (*Uint64)(nil)

func Uint64FromStd(v *atomic.Uint64) *Uint64 {
	return (*Uint64)(unsafe.Pointer(v))
}

func Uint64FromPointer(v *uint64) *Uint64 {
	return (*Uint64)(unsafe.Pointer(v))
}

func (v *Uint64) Generic() *Integer[uint64] {
	return (*Integer[uint64])(v)
}

func (v *Uint64) Std() *atomic.Uint64 {
	return (*atomic.Uint64)(unsafe.Pointer(v))
}

func (v *Uint64) Load() uint64 {
	return atomic.LoadUint64((*uint64)(&v.Value))
}

func (v *Uint64) Store(new uint64) {
	atomic.StoreUint64((*uint64)(&v.Value), new)
}

func (v *Uint64) Add(delta uint64) uint64 {
	return atomic.AddUint64((*uint64)(&v.Value), delta)
}

func (v *Uint64) Swap(new uint64) uint64 {
	return atomic.SwapUint64((*uint64)(&v.Value), new)
}

func (v *Uint64) CompareAndSwap(old, new uint64) bool {
	return atomic.CompareAndSwapUint64((*uint64)(&v.Value), old, new)
}
