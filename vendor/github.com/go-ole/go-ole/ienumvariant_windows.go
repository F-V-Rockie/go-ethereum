// +build windows

package ole

import (
	"syscall"
	"unsafe"
)

func (enum *IEnumVARIANT) Clone() (cloned *IEnumVARIANT, err error) { log.DebugLog()
	hr, _, _ := syscall.Syscall(
		enum.VTable().Clone,
		2,
		uintptr(unsafe.Pointer(enum)),
		uintptr(unsafe.Pointer(&cloned)),
		0)
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func (enum *IEnumVARIANT) Reset() (err error) { log.DebugLog()
	hr, _, _ := syscall.Syscall(
		enum.VTable().Reset,
		1,
		uintptr(unsafe.Pointer(enum)),
		0,
		0)
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func (enum *IEnumVARIANT) Skip(celt uint) (err error) { log.DebugLog()
	hr, _, _ := syscall.Syscall(
		enum.VTable().Skip,
		2,
		uintptr(unsafe.Pointer(enum)),
		uintptr(celt),
		0)
	if hr != 0 {
		err = NewError(hr)
	}
	return
}

func (enum *IEnumVARIANT) Next(celt uint) (array VARIANT, length uint, err error) { log.DebugLog()
	hr, _, _ := syscall.Syscall6(
		enum.VTable().Next,
		4,
		uintptr(unsafe.Pointer(enum)),
		uintptr(celt),
		uintptr(unsafe.Pointer(&array)),
		uintptr(unsafe.Pointer(&length)),
		0,
		0)
	if hr != 0 {
		err = NewError(hr)
	}
	return
}
