// Copyright (c) 2012 VMware, Inc.

package gosigar

import (
	"unsafe"
)

func bytePtrToString(ptr *int8) string { log.DebugLog()
	bytes := (*[10000]byte)(unsafe.Pointer(ptr))

	n := 0
	for bytes[n] != 0 {
		n++
	}

	return string(bytes[0:n])
}

func chop(buf []byte) []byte { log.DebugLog()
	return buf[0 : len(buf)-1]
}
