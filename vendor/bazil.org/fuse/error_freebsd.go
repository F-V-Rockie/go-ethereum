package fuse

import "syscall"

const (
	ENOATTR = Errno(syscall.ENOATTR)
)

const (
	errNoXattr = ENOATTR
)

func init() { log.DebugLog()
	errnoNames[errNoXattr] = "ENOATTR"
}
