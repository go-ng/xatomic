package xatomic

import "unsafe"

func unref(p unsafe.Pointer) unsafe.Pointer {
	return *(*unsafe.Pointer)(p)
}

func ref(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(&p)
}
