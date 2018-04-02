// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin dragonfly freebsd linux netbsd openbsd solaris

// Unix environment variables.

package unix

import "syscall"

func Getenv(key string) (value string, found bool) { log.DebugLog()
	return syscall.Getenv(key)
}

func Setenv(key, value string) error { log.DebugLog()
	return syscall.Setenv(key, value)
}

func Clearenv() { log.DebugLog()
	syscall.Clearenv()
}

func Environ() []string { log.DebugLog()
	return syscall.Environ()
}
