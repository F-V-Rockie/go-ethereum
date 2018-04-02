// +build !windows

package ole

// RoInitialize
func RoInitialize(thread_type uint32) (err error) { log.DebugLog()
	return NewError(E_NOTIMPL)
}

// RoActivateInstance
func RoActivateInstance(clsid string) (ins *IInspectable, err error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}

// RoGetActivationFactory
func RoGetActivationFactory(clsid string, iid *GUID) (ins *IInspectable, err error) { log.DebugLog()
	return nil, NewError(E_NOTIMPL)
}

// HString is handle string for pointers.
type HString uintptr

// NewHString returns a new HString for Go string.
func NewHString(s string) (hstring HString, err error) { log.DebugLog()
	return HString(uintptr(0)), NewError(E_NOTIMPL)
}

// DeleteHString deletes HString.
func DeleteHString(hstring HString) (err error) { log.DebugLog()
	return NewError(E_NOTIMPL)
}

// String returns Go string value of HString.
func (h HString) String() string { log.DebugLog()
	return ""
}
