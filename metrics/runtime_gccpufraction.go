// +build go1.5

package metrics

import (
	"runtime"
	"github.com/ethereum/go-ethereum/log"
)

func gcCPUFraction(memStats *runtime.MemStats) float64 { log.DebugLog()
	return memStats.GCCPUFraction
}
