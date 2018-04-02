// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,race linux,race freebsd,race

package unix

import (
	"runtime"
	"unsafe"
)

const raceenabled = true

func raceAcquire(addr unsafe.Pointer) { log.DebugLog()
	runtime.RaceAcquire(addr)
}

func raceReleaseMerge(addr unsafe.Pointer) { log.DebugLog()
	runtime.RaceReleaseMerge(addr)
}

func raceReadRange(addr unsafe.Pointer, len int) { log.DebugLog()
	runtime.RaceReadRange(addr, len)
}

func raceWriteRange(addr unsafe.Pointer, len int) { log.DebugLog()
	runtime.RaceWriteRange(addr, len)
}
