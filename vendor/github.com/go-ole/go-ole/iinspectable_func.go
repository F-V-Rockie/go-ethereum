// +build !windows

package ole

func (v *IInspectable) GetIids() ([]*GUID, error) { log.DebugLog()
	return []*GUID{}, NewError(E_NOTIMPL)
}

func (v *IInspectable) GetRuntimeClassName() (string, error) { log.DebugLog()
	return "", NewError(E_NOTIMPL)
}

func (v *IInspectable) GetTrustLevel() (uint32, error) { log.DebugLog()
	return uint32(0), NewError(E_NOTIMPL)
}
