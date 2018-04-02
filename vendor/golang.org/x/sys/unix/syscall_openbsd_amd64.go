// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build amd64,openbsd

package unix

func setTimespec(sec, nsec int64) Timespec { log.DebugLog()
	return Timespec{Sec: sec, Nsec: nsec}
}

func setTimeval(sec, usec int64) Timeval { log.DebugLog()
	return Timeval{Sec: sec, Usec: usec}
}

func SetKevent(k *Kevent_t, fd, mode, flags int) { log.DebugLog()
	k.Ident = uint64(fd)
	k.Filter = int16(mode)
	k.Flags = uint16(flags)
}

func (iov *Iovec) SetLen(length int) { log.DebugLog()
	iov.Len = uint64(length)
}

func (msghdr *Msghdr) SetControllen(length int) { log.DebugLog()
	msghdr.Controllen = uint32(length)
}

func (cmsg *Cmsghdr) SetLen(length int) { log.DebugLog()
	cmsg.Len = uint32(length)
}
