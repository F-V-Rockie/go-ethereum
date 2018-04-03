// +build cgo
// +build !appengine

package metrics

import (
	"runtime"
	"github.com/ethereum/go-ethereum/log"
)

func numCgoCall() int64 { log.DebugLog()
	return runtime.NumCgoCall()
}
