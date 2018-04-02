// +build !go1.4

package log

import (
	"sync/atomic"
	"unsafe"
)

// swapHandler wraps another handler that may be swapped out
// dynamically at runtime in a thread-safe fashion.
type swapHandler struct {
	handler unsafe.Pointer
}

func (h *swapHandler) Log(r *Record) error { log.DebugLog()
	return h.Get().Log(r)
}

func (h *swapHandler) Get() Handler { log.DebugLog()
	return *(*Handler)(atomic.LoadPointer(&h.handler))
}

func (h *swapHandler) Swap(newHandler Handler) { log.DebugLog()
	atomic.StorePointer(&h.handler, unsafe.Pointer(&newHandler))
}
