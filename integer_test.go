package xatomic

import (
	"sync/atomic"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestIntegerConstructor(t *testing.T) {
	require.Equal(t, unsafe.Sizeof(int32(0)), unsafe.Sizeof(Integer[int32]{}))
	require.Equal(t, unsafe.Sizeof(int64(0)), unsafe.Sizeof(Integer[int64]{}))
	require.Equal(t, unsafe.Sizeof(uint32(0)), unsafe.Sizeof(Integer[uint32]{}))
	require.Equal(t, unsafe.Sizeof(uint64(0)), unsafe.Sizeof(Integer[uint64]{}))
	require.Equal(t, unsafe.Sizeof(atomic.Int32{}), unsafe.Sizeof(Integer[int32]{}))
	require.Equal(t, unsafe.Sizeof(atomic.Int64{}), unsafe.Sizeof(Integer[int64]{}))
	require.Equal(t, unsafe.Sizeof(atomic.Uint32{}), unsafe.Sizeof(Integer[uint32]{}))
	require.Equal(t, unsafe.Sizeof(atomic.Uint64{}), unsafe.Sizeof(Integer[uint64]{}))

	var std atomic.Int32
	std.Store(0x12345678)
	require.Equal(t, int32(0x12345678), Int32FromStd(&std).Load())

	var v int32 = 0x12345678
	require.Equal(t, int32(0x12345678), Int32FromPointer(&v).Load())
}
