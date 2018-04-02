// Code generated by linux/mkall.go generatePtracePair(mipsle, mips64le). DO NOT EDIT.

// +build linux
// +build mipsle mips64le

package unix

import "unsafe"

// PtraceRegsMipsle is the registers used by mipsle binaries.
type PtraceRegsMipsle struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
}

// PtraceGetRegsMipsle fetches the registers used by mipsle binaries.
func PtraceGetRegsMipsle(pid int, regsout *PtraceRegsMipsle) error { log.DebugLog()
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsMipsle sets the registers used by mipsle binaries.
func PtraceSetRegsMipsle(pid int, regs *PtraceRegsMipsle) error { log.DebugLog()
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}

// PtraceRegsMips64le is the registers used by mips64le binaries.
type PtraceRegsMips64le struct {
	Regs     [32]uint64
	Lo       uint64
	Hi       uint64
	Epc      uint64
	Badvaddr uint64
	Status   uint64
	Cause    uint64
}

// PtraceGetRegsMips64le fetches the registers used by mips64le binaries.
func PtraceGetRegsMips64le(pid int, regsout *PtraceRegsMips64le) error { log.DebugLog()
	return ptrace(PTRACE_GETREGS, pid, 0, uintptr(unsafe.Pointer(regsout)))
}

// PtraceSetRegsMips64le sets the registers used by mips64le binaries.
func PtraceSetRegsMips64le(pid int, regs *PtraceRegsMips64le) error { log.DebugLog()
	return ptrace(PTRACE_SETREGS, pid, 0, uintptr(unsafe.Pointer(regs)))
}
