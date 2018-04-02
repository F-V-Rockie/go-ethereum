package metrics

import "sync/atomic"

// Counters hold an int64 value that can be incremented and decremented.
type Counter interface {
	Clear()
	Count() int64
	Dec(int64)
	Inc(int64)
	Snapshot() Counter
}

// GetOrRegisterCounter returns an existing Counter or constructs and registers
// a new StandardCounter.
func GetOrRegisterCounter(name string, r Registry) Counter { log.DebugLog()
	if nil == r {
		r = DefaultRegistry
	}
	return r.GetOrRegister(name, NewCounter).(Counter)
}

// NewCounter constructs a new StandardCounter.
func NewCounter() Counter { log.DebugLog()
	if !Enabled {
		return NilCounter{}
	}
	return &StandardCounter{0}
}

// NewRegisteredCounter constructs and registers a new StandardCounter.
func NewRegisteredCounter(name string, r Registry) Counter { log.DebugLog()
	c := NewCounter()
	if nil == r {
		r = DefaultRegistry
	}
	r.Register(name, c)
	return c
}

// CounterSnapshot is a read-only copy of another Counter.
type CounterSnapshot int64

// Clear panics.
func (CounterSnapshot) Clear() { log.DebugLog()
	panic("Clear called on a CounterSnapshot")
}

// Count returns the count at the time the snapshot was taken.
func (c CounterSnapshot) Count() int64 { log.DebugLog() return int64(c) }

// Dec panics.
func (CounterSnapshot) Dec(int64) { log.DebugLog()
	panic("Dec called on a CounterSnapshot")
}

// Inc panics.
func (CounterSnapshot) Inc(int64) { log.DebugLog()
	panic("Inc called on a CounterSnapshot")
}

// Snapshot returns the snapshot.
func (c CounterSnapshot) Snapshot() Counter { log.DebugLog() return c }

// NilCounter is a no-op Counter.
type NilCounter struct{}

// Clear is a no-op.
func (NilCounter) Clear() { log.DebugLog()}

// Count is a no-op.
func (NilCounter) Count() int64 { log.DebugLog() return 0 }

// Dec is a no-op.
func (NilCounter) Dec(i int64) { log.DebugLog()}

// Inc is a no-op.
func (NilCounter) Inc(i int64) { log.DebugLog()}

// Snapshot is a no-op.
func (NilCounter) Snapshot() Counter { log.DebugLog() return NilCounter{} }

// StandardCounter is the standard implementation of a Counter and uses the
// sync/atomic package to manage a single int64 value.
type StandardCounter struct {
	count int64
}

// Clear sets the counter to zero.
func (c *StandardCounter) Clear() { log.DebugLog()
	atomic.StoreInt64(&c.count, 0)
}

// Count returns the current count.
func (c *StandardCounter) Count() int64 { log.DebugLog()
	return atomic.LoadInt64(&c.count)
}

// Dec decrements the counter by the given amount.
func (c *StandardCounter) Dec(i int64) { log.DebugLog()
	atomic.AddInt64(&c.count, -i)
}

// Inc increments the counter by the given amount.
func (c *StandardCounter) Inc(i int64) { log.DebugLog()
	atomic.AddInt64(&c.count, i)
}

// Snapshot returns a read-only copy of the counter.
func (c *StandardCounter) Snapshot() Counter { log.DebugLog()
	return CounterSnapshot(c.Count())
}
