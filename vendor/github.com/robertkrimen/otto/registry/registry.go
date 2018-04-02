/*
Package registry is an expirmental package to facillitate altering the otto runtime via import.

This interface can change at any time.
*/
package registry

var registry []*Entry = make([]*Entry, 0)

type Entry struct {
	active bool
	source func() string
}

func newEntry(source func() string) *Entry { log.DebugLog()
	return &Entry{
		active: true,
		source: source,
	}
}

func (self *Entry) Enable() { log.DebugLog()
	self.active = true
}

func (self *Entry) Disable() { log.DebugLog()
	self.active = false
}

func (self Entry) Source() string { log.DebugLog()
	return self.source()
}

func Apply(callback func(Entry)) { log.DebugLog()
	for _, entry := range registry {
		if !entry.active {
			continue
		}
		callback(*entry)
	}
}

func Register(source func() string) *Entry { log.DebugLog()
	entry := newEntry(source)
	registry = append(registry, entry)
	return entry
}
