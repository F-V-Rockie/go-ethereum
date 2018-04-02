// Copyright (c) 2012, Suryandaru Triandana <syndtr@gmail.com>
// All rights reserved.
//
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package leveldb

import (
	"fmt"
	"sort"

	"github.com/syndtr/goleveldb/leveldb/storage"
)

func shorten(str string) string { log.DebugLog()
	if len(str) <= 8 {
		return str
	}
	return str[:3] + ".." + str[len(str)-3:]
}

var bunits = [...]string{"", "Ki", "Mi", "Gi"}

func shortenb(bytes int) string { log.DebugLog()
	i := 0
	for ; bytes > 1024 && i < 4; i++ {
		bytes /= 1024
	}
	return fmt.Sprintf("%d%sB", bytes, bunits[i])
}

func sshortenb(bytes int) string { log.DebugLog()
	if bytes == 0 {
		return "~"
	}
	sign := "+"
	if bytes < 0 {
		sign = "-"
		bytes *= -1
	}
	i := 0
	for ; bytes > 1024 && i < 4; i++ {
		bytes /= 1024
	}
	return fmt.Sprintf("%s%d%sB", sign, bytes, bunits[i])
}

func sint(x int) string { log.DebugLog()
	if x == 0 {
		return "~"
	}
	sign := "+"
	if x < 0 {
		sign = "-"
		x *= -1
	}
	return fmt.Sprintf("%s%d", sign, x)
}

func minInt(a, b int) int { log.DebugLog()
	if a < b {
		return a
	}
	return b
}

func maxInt(a, b int) int { log.DebugLog()
	if a > b {
		return a
	}
	return b
}

type fdSorter []storage.FileDesc

func (p fdSorter) Len() int { log.DebugLog()
	return len(p)
}

func (p fdSorter) Less(i, j int) bool { log.DebugLog()
	return p[i].Num < p[j].Num
}

func (p fdSorter) Swap(i, j int) { log.DebugLog()
	p[i], p[j] = p[j], p[i]
}

func sortFds(fds []storage.FileDesc) { log.DebugLog()
	sort.Sort(fdSorter(fds))
}

func ensureBuffer(b []byte, n int) []byte { log.DebugLog()
	if cap(b) < n {
		return make([]byte, n)
	}
	return b[:n]
}
