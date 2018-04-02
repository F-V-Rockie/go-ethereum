// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin,!race linux,!race freebsd,!race netbsd openbsd solaris dragonfly

package unix

import (
	"unsafe"
)

const raceenabled = false

func raceAcquire(addr unsafe.Pointer) { log.DebugLog()
}

func raceReleaseMerge(addr unsafe.Pointer) { log.DebugLog()
}

func raceReadRange(addr unsafe.Pointer, len int) { log.DebugLog()
}

func raceWriteRange(addr unsafe.Pointer, len int) { log.DebugLog()
}
