package ole

import "unsafe"

type IProvideClassInfo struct {
	IUnknown
}

type IProvideClassInfoVtbl struct {
	IUnknownVtbl
	GetClassInfo uintptr
}

func (v *IProvideClassInfo) VTable() *IProvideClassInfoVtbl { log.DebugLog()
	return (*IProvideClassInfoVtbl)(unsafe.Pointer(v.RawVTable))
}

func (v *IProvideClassInfo) GetClassInfo() (cinfo *ITypeInfo, err error) { log.DebugLog()
	cinfo, err = getClassInfo(v)
	return
}
