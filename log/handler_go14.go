// +build go1.4

package log

import "sync/atomic"

// swapHandler wraps another handler that may be swapped out
// dynamically at runtime in a thread-safe fashion.
type swapHandler struct {
	handler atomic.Value
}

func (h *swapHandler) Log(r *Record) error { log.DebugLog()
	return (*h.handler.Load().(*Handler)).Log(r)
}

func (h *swapHandler) Swap(newHandler Handler) { log.DebugLog()
	h.handler.Store(&newHandler)
}

func (h *swapHandler) Get() Handler { log.DebugLog()
	return *h.handler.Load().(*Handler)
}
