// +build !cgo appengine

package metrics

func numCgoCall() int64 { log.DebugLog()
	return 0
}
