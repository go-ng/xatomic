package xatomic

import (
	"sync/atomic"
	"unsafe"
)

type Uint32 Integer[uint32]

var _ Abstract[uint32] = (*Uint32)(nil)

func Uint32FromStd(v *atomic.Uint32) *Uint32 {
	return (*Uint32)(unsafe.Pointer(v))
}

func Uint32FromPointer(v *uint32) *Uint32 {
	return (*Uint32)(unsafe.Pointer(v))
}

func (v *Uint32) Load() uint32 {
	return atomic.LoadUint32((*uint32)(&v.Value))
}

func (v *Uint32) Store(new uint32) {
	atomic.StoreUint32((*uint32)(&v.Value), new)
}

func (v *Uint32) Add(delta uint32) uint32 {
	return atomic.AddUint32((*uint32)(&v.Value), delta)
}

func (v *Uint32) Swap(new uint32) uint32 {
	return atomic.SwapUint32((*uint32)(&v.Value), new)
}

func (v *Uint32) CompareAndSwap(old, new uint32) bool {
	return atomic.CompareAndSwapUint32((*uint32)(&v.Value), old, new)
}
