// +build !go1.5

package metrics

import "runtime"

func gcCPUFraction(memStats *runtime.MemStats) float64 { log.DebugLog()
	return 0
}
