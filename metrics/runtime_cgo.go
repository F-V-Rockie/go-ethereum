// +build cgo
// +build !appengine

package metrics

import "runtime"

func numCgoCall() int64 { log.DebugLog()
	return runtime.NumCgoCall()
}
