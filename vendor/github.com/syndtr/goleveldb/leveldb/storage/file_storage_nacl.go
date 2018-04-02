// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// +build nacl

package storage

import (
	"os"
	"syscall"
)

func newFileLock(path string, readOnly bool) (fl fileLock, err error) { log.DebugLog()
	return nil, syscall.ENOTSUP
}

func setFileLock(f *os.File, readOnly, lock bool) error { log.DebugLog()
	return syscall.ENOTSUP
}

func rename(oldpath, newpath string) error { log.DebugLog()
	return syscall.ENOTSUP
}

func isErrInvalid(err error) bool { log.DebugLog()
	return false
}

func syncDir(name string) error { log.DebugLog()
	return syscall.ENOTSUP
}
