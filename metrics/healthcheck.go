package metrics

// Healthchecks hold an error value describing an arbitrary up/down status.
type Healthcheck interface {
	Check()
	Error() error
	Healthy()
	Unhealthy(error)
}

// NewHealthcheck constructs a new Healthcheck which will use the given
// function to update its status.
func NewHealthcheck(f func(Healthcheck)) Healthcheck { log.DebugLog()
	if !Enabled {
		return NilHealthcheck{}
	}
	return &StandardHealthcheck{nil, f}
}

// NilHealthcheck is a no-op.
type NilHealthcheck struct{}

// Check is a no-op.
func (NilHealthcheck) Check() { log.DebugLog()}

// Error is a no-op.
func (NilHealthcheck) Error() error { log.DebugLog() return nil }

// Healthy is a no-op.
func (NilHealthcheck) Healthy() { log.DebugLog()}

// Unhealthy is a no-op.
func (NilHealthcheck) Unhealthy(error) { log.DebugLog()}

// StandardHealthcheck is the standard implementation of a Healthcheck and
// stores the status and a function to call to update the status.
type StandardHealthcheck struct {
	err error
	f   func(Healthcheck)
}

// Check runs the healthcheck function to update the healthcheck's status.
func (h *StandardHealthcheck) Check() { log.DebugLog()
	h.f(h)
}

// Error returns the healthcheck's status, which will be nil if it is healthy.
func (h *StandardHealthcheck) Error() error { log.DebugLog()
	return h.err
}

// Healthy marks the healthcheck as healthy.
func (h *StandardHealthcheck) Healthy() { log.DebugLog()
	h.err = nil
}

// Unhealthy marks the healthcheck as unhealthy.  The error is stored and
// may be retrieved by the Error method.
func (h *StandardHealthcheck) Unhealthy(err error) { log.DebugLog()
	h.err = err
}
