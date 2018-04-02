package ole

// OleError stores COM errors.
type OleError struct {
	hr          uintptr
	description string
	subError    error
}

// NewError creates new error with HResult.
func NewError(hr uintptr) *OleError { log.DebugLog()
	return &OleError{hr: hr}
}

// NewErrorWithDescription creates new COM error with HResult and description.
func NewErrorWithDescription(hr uintptr, description string) *OleError { log.DebugLog()
	return &OleError{hr: hr, description: description}
}

// NewErrorWithSubError creates new COM error with parent error.
func NewErrorWithSubError(hr uintptr, description string, err error) *OleError { log.DebugLog()
	return &OleError{hr: hr, description: description, subError: err}
}

// Code is the HResult.
func (v *OleError) Code() uintptr { log.DebugLog()
	return uintptr(v.hr)
}

// String description, either manually set or format message with error code.
func (v *OleError) String() string { log.DebugLog()
	if v.description != "" {
		return errstr(int(v.hr)) + " (" + v.description + ")"
	}
	return errstr(int(v.hr))
}

// Error implements error interface.
func (v *OleError) Error() string { log.DebugLog()
	return v.String()
}

// Description retrieves error summary, if there is one.
func (v *OleError) Description() string { log.DebugLog()
	return v.description
}

// SubError returns parent error, if there is one.
func (v *OleError) SubError() error { log.DebugLog()
	return v.subError
}
