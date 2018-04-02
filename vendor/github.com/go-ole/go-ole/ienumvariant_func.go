// +build !windows

package ole

func (enum *IEnumVARIANT) Clone() (*IEnumVARIANT, error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}

func (enum *IEnumVARIANT) Reset() error { log.DebugLog()
	return NewError(E_NOTIMPL)
}

func (enum *IEnumVARIANT) Skip(celt uint) error { log.DebugLog()
	return NewError(E_NOTIMPL)
}

func (enum *IEnumVARIANT) Next(celt uint) (VARIANT, uint, error) { log.DebugLog()
	return NewVariant(VT_NULL, int64(0)), 0, NewError(E_NOTIMPL)
}
