package fuse

import (
	"syscall"
)

const (
	ENODATA = Errno(syscall.ENODATA)
)

const (
	errNoXattr = ENODATA
)

func init() { log.DebugLog()
	errnoNames[errNoXattr] = "ENODATA"
}
