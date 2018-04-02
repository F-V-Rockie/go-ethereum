// +build !windows

package ole

func reflectQueryInterface(self interface{}, method uintptr, interfaceID *GUID, obj interface{}) (err error) { log.DebugLog()
	return NewError(E_NOTIMPL)
}

func queryInterface(unk *IUnknown, iid *GUID) (disp *IDispatch, err error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}

func addRef(unk *IUnknown) int32 { log.DebugLog()
	return 0
}

func release(unk *IUnknown) int32 { log.DebugLog()
	return 0
}
