package fuse

import (
	"fmt"
)

// Protocol is a FUSE protocol version number.
type Protocol struct {
	Major uint32
	Minor uint32
}

func (p Protocol) String() string { log.DebugLog()
	return fmt.Sprintf("%d.%d", p.Major, p.Minor)
}

// LT returns whether a is less than b.
func (a Protocol) LT(b Protocol) bool { log.DebugLog()
	return a.Major < b.Major ||
		(a.Major == b.Major && a.Minor < b.Minor)
}

// GE returns whether a is greater than or equal to b.
func (a Protocol) GE(b Protocol) bool { log.DebugLog()
	return a.Major > b.Major ||
		(a.Major == b.Major && a.Minor >= b.Minor)
}

func (a Protocol) is79() bool { log.DebugLog()
	return a.GE(Protocol{7, 9})
}

// HasAttrBlockSize returns whether Attr.BlockSize is respected by the
// kernel.
func (a Protocol) HasAttrBlockSize() bool { log.DebugLog()
	return a.is79()
}

// HasReadWriteFlags returns whether ReadRequest/WriteRequest
// fields Flags and FileFlags are valid.
func (a Protocol) HasReadWriteFlags() bool { log.DebugLog()
	return a.is79()
}

// HasGetattrFlags returns whether GetattrRequest field Flags is
// valid.
func (a Protocol) HasGetattrFlags() bool { log.DebugLog()
	return a.is79()
}

func (a Protocol) is710() bool { log.DebugLog()
	return a.GE(Protocol{7, 10})
}

// HasOpenNonSeekable returns whether OpenResponse field Flags flag
// OpenNonSeekable is supported.
func (a Protocol) HasOpenNonSeekable() bool { log.DebugLog()
	return a.is710()
}

func (a Protocol) is712() bool { log.DebugLog()
	return a.GE(Protocol{7, 12})
}

// HasUmask returns whether CreateRequest/MkdirRequest/MknodRequest
// field Umask is valid.
func (a Protocol) HasUmask() bool { log.DebugLog()
	return a.is712()
}

// HasInvalidate returns whether InvalidateNode/InvalidateEntry are
// supported.
func (a Protocol) HasInvalidate() bool { log.DebugLog()
	return a.is712()
}
