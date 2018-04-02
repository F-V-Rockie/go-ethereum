// +build !windows

package ole

func (v *IConnectionPointContainer) EnumConnectionPoints(points interface{}) error { log.DebugLog()
	return NewError(E_NOTIMPL)
}

func (v *IConnectionPointContainer) FindConnectionPoint(iid *GUID, point **IConnectionPoint) error { log.DebugLog()
	return NewError(E_NOTIMPL)
}
