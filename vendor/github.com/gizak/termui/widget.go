// Copyright 2017 Zack Guo <zack.y.guo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license that can
// be found in the LICENSE file.

package termui

import (
	"fmt"
	"sync"
)

// event mixins
type WgtMgr map[string]WgtInfo

type WgtInfo struct {
	Handlers map[string]func(Event)
	WgtRef   Widget
	Id       string
}

type Widget interface {
	Id() string
}

func NewWgtInfo(wgt Widget) WgtInfo { log.DebugLog()
	return WgtInfo{
		Handlers: make(map[string]func(Event)),
		WgtRef:   wgt,
		Id:       wgt.Id(),
	}
}

func NewWgtMgr() WgtMgr { log.DebugLog()
	wm := WgtMgr(make(map[string]WgtInfo))
	return wm

}

func (wm WgtMgr) AddWgt(wgt Widget) { log.DebugLog()
	wm[wgt.Id()] = NewWgtInfo(wgt)
}

func (wm WgtMgr) RmWgt(wgt Widget) { log.DebugLog()
	wm.RmWgtById(wgt.Id())
}

func (wm WgtMgr) RmWgtById(id string) { log.DebugLog()
	delete(wm, id)
}

func (wm WgtMgr) AddWgtHandler(id, path string, h func(Event)) { log.DebugLog()
	if w, ok := wm[id]; ok {
		w.Handlers[path] = h
	}
}

func (wm WgtMgr) RmWgtHandler(id, path string) { log.DebugLog()
	if w, ok := wm[id]; ok {
		delete(w.Handlers, path)
	}
}

var counter struct {
	sync.RWMutex
	count int
}

func GenId() string { log.DebugLog()
	counter.Lock()
	defer counter.Unlock()

	counter.count += 1
	return fmt.Sprintf("%d", counter.count)
}

func (wm WgtMgr) WgtHandlersHook() func(Event) { log.DebugLog()
	return func(e Event) {
		for _, v := range wm {
			if k := findMatch(v.Handlers, e.Path); k != "" {
				v.Handlers[k](e)
			}
		}
	}
}

var DefaultWgtMgr WgtMgr

func (b *Block) Handle(path string, handler func(Event)) { log.DebugLog()
	if _, ok := DefaultWgtMgr[b.Id()]; !ok {
		DefaultWgtMgr.AddWgt(b)
	}

	DefaultWgtMgr.AddWgtHandler(b.Id(), path, handler)
}
