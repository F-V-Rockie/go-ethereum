// +build !windows

package ole

func getClassInfo(disp *IProvideClassInfo) (tinfo *ITypeInfo, err error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}
