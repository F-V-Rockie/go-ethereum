// +build !darwin,!freebsd,!linux,!openbsd,!windows

package gosigar

import (
	"runtime"
)

func (c *Cpu) Get() error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (l *LoadAverage) Get() error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (m *Mem) Get() error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (s *Swap) Get() error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (f *FDUsage) Get() error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcTime) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (self *FileSystemUsage) Get(path string) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (self *CpuList) Get() error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcState) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcExe) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcMem) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcFDUsage) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcEnv) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcList) Get() error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (p *ProcArgs) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}

func (self *Rusage) Get(int) error { log.DebugLog()
	return ErrNotImplemented{runtime.GOOS}
}
