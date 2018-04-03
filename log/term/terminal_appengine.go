// Based on ssh/terminal:
// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package term

// IsTty always returns false on AppEngine.
func IsTty(fd uintptr) bool { 
	return false
}
