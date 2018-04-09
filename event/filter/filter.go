// Copyright 2014 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package filter implements event filters.
package filter

import "reflect"
import "github.com/ethereum/go-ethereum/log"

type Filter interface {
	Compare(Filter) bool
	Trigger(data interface{})
}

type FilterEvent struct {
	filter Filter
	data   interface{}
}

type Filters struct {
	id       int
	watchers map[int]Filter
	ch       chan FilterEvent

	quit chan struct{}
}

func New() *Filters { log.DebugLog()
	return &Filters{
		ch:       make(chan FilterEvent),
		watchers: make(map[int]Filter),
		quit:     make(chan struct{}),
	}
}

func (self *Filters) Start() { log.DebugLog()
	go self.loop()
}

func (self *Filters) Stop() { log.DebugLog()
	close(self.quit)
}

func (self *Filters) Notify(filter Filter, data interface{}) { log.DebugLog()
	self.ch <- FilterEvent{filter, data}
}

func (self *Filters) Install(watcher Filter) int { log.DebugLog()
	self.watchers[self.id] = watcher
	self.id++

	return self.id - 1
}

func (self *Filters) Uninstall(id int) { log.DebugLog()
	delete(self.watchers, id)
}

func (self *Filters) loop() { log.DebugLog()
out:
	for {
		select {
		case <-self.quit:
			break out
		case event := <-self.ch:
			for _, watcher := range self.watchers {
				if reflect.TypeOf(watcher) == reflect.TypeOf(event.filter) {
					if watcher.Compare(event.filter) {
						watcher.Trigger(event.data)
					}
				}
			}
		}
	}
}

func (self *Filters) Match(a, b Filter) bool { log.DebugLog()
	return reflect.TypeOf(a) == reflect.TypeOf(b) && a.Compare(b)
}

func (self *Filters) Get(i int) Filter { log.DebugLog()
	return self.watchers[i]
}
