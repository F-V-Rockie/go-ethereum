// +build !windows

package ole

func (v *ITypeInfo) GetTypeAttr() (*TYPEATTR, error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}
