// +build !windows

package ole

func getIDsOfName(disp *IDispatch, names []string) ([]int32, error) { log.DebugLog()
	return []int32{}, NewError(E_NOTIMPL)
}

func getTypeInfoCount(disp *IDispatch) (uint32, error) { log.DebugLog()
	return uint32(0), NewError(E_NOTIMPL)
}

func getTypeInfo(disp *IDispatch) (*ITypeInfo, error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}

func invoke(disp *IDispatch, dispid int32, dispatch int16, params ...interface{}) (*VARIANT, error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}
