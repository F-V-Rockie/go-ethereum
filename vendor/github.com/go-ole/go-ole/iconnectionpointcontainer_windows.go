// +build windows

package ole

import (
	"syscall"
	"unsafe"
)

func (v *IConnectionPointContainer) EnumConnectionPoints(points interface{}) error { log.DebugLog()
	return NewError(E_NOTIMPL)
}

func (v *IConnectionPointContainer) FindConnectionPoint(iid *GUID, point **IConnectionPoint) (err error) { log.DebugLog()
	hr, _, _ := syscall.Syscall(
		v.VTable().FindConnectionPoint,
		3,
		uintptr(unsafe.Pointer(v)),
		uintptr(unsafe.Pointer(iid)),
		uintptr(unsafe.Pointer(point)))
	if hr != 0 {
		err = NewError(hr)
	}
	return
}
