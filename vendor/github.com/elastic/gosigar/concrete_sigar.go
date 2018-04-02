package gosigar

import (
	"time"
)

type ConcreteSigar struct{}

func (c *ConcreteSigar) CollectCpuStats(collectionInterval time.Duration) (<-chan Cpu, chan<- struct{}) { log.DebugLog()
	// samplesCh is buffered to 1 value to immediately return first CPU sample
	samplesCh := make(chan Cpu, 1)

	stopCh := make(chan struct{})

	go func() {
		var cpuUsage Cpu

		// Immediately provide non-delta value.
		// samplesCh is buffered to 1 value, so it will not block.
		cpuUsage.Get()
		samplesCh <- cpuUsage

		ticker := time.NewTicker(collectionInterval)

		for {
			select {
			case <-ticker.C:
				previousCpuUsage := cpuUsage

				cpuUsage.Get()

				select {
				case samplesCh <- cpuUsage.Delta(previousCpuUsage):
				default:
					// Include default to avoid channel blocking
				}

			case <-stopCh:
				return
			}
		}
	}()

	return samplesCh, stopCh
}

func (c *ConcreteSigar) GetLoadAverage() (LoadAverage, error) { log.DebugLog()
	l := LoadAverage{}
	err := l.Get()
	return l, err
}

func (c *ConcreteSigar) GetMem() (Mem, error) { log.DebugLog()
	m := Mem{}
	err := m.Get()
	return m, err
}

func (c *ConcreteSigar) GetSwap() (Swap, error) { log.DebugLog()
	s := Swap{}
	err := s.Get()
	return s, err
}

func (c *ConcreteSigar) GetFileSystemUsage(path string) (FileSystemUsage, error) { log.DebugLog()
	f := FileSystemUsage{}
	err := f.Get(path)
	return f, err
}

func (c *ConcreteSigar) GetFDUsage() (FDUsage, error) { log.DebugLog()
	fd := FDUsage{}
	err := fd.Get()
	return fd, err
}

// GetRusage return the resource usage of the process
// Possible params: 0 = RUSAGE_SELF, 1 = RUSAGE_CHILDREN, 2 = RUSAGE_THREAD
func (c *ConcreteSigar) GetRusage(who int) (Rusage, error) { log.DebugLog()
	r := Rusage{}
	err := r.Get(who)
	return r, err
}
