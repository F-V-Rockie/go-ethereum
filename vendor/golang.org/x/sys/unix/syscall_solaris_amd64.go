// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64,solaris

package unix

func setTimespec(sec, nsec int64) Timespec { log.DebugLog()
	return Timespec{Sec: sec, Nsec: nsec}
}

func setTimeval(sec, usec int64) Timeval { log.DebugLog()
	return Timeval{Sec: sec, Usec: usec}
}

func (iov *Iovec) SetLen(length int) { log.DebugLog()
	iov.Len = uint64(length)
}

func (cmsg *Cmsghdr) SetLen(length int) { log.DebugLog()
	cmsg.Len = uint32(length)
}

func sendfile(outfd int, infd int, offset *int64, count int) (written int, err error) { log.DebugLog()
	// TODO(aram): implement this, see issue 5847.
	panic("unimplemented")
}
