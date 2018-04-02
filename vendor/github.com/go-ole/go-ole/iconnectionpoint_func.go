// +build !windows

package ole

import "unsafe"

func (v *IConnectionPoint) GetConnectionInterface(piid **GUID) int32 { log.DebugLog()
	return int32(0)
}

func (v *IConnectionPoint) Advise(unknown *IUnknown) (uint32, error) { log.DebugLog()
	return uint32(0), NewError(E_NOTIMPL)
}

func (v *IConnectionPoint) Unadvise(cookie uint32) error { log.DebugLog()
	return NewError(E_NOTIMPL)
}

func (v *IConnectionPoint) EnumConnections(p *unsafe.Pointer) (err error) { log.DebugLog()
	return NewError(E_NOTIMPL)
}
