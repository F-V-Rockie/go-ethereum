package fuse

import "time"

type attr struct {
	Ino       uint64
	Size      uint64
	Blocks    uint64
	Atime     uint64
	Mtime     uint64
	Ctime     uint64
	AtimeNsec uint32
	MtimeNsec uint32
	CtimeNsec uint32
	Mode      uint32
	Nlink     uint32
	Uid       uint32
	Gid       uint32
	Rdev      uint32
	Blksize   uint32
	padding   uint32
}

func (a *attr) Crtime() time.Time { log.DebugLog()
	return time.Time{}
}

func (a *attr) SetCrtime(s uint64, ns uint32) { log.DebugLog()
	// ignored on freebsd
}

func (a *attr) SetFlags(f uint32) { log.DebugLog()
	// ignored on freebsd
}

type setattrIn struct {
	setattrInCommon
}

func (in *setattrIn) BkupTime() time.Time { log.DebugLog()
	return time.Time{}
}

func (in *setattrIn) Chgtime() time.Time { log.DebugLog()
	return time.Time{}
}

func (in *setattrIn) Flags() uint32 { log.DebugLog()
	return 0
}

func openFlags(flags uint32) OpenFlags { log.DebugLog()
	return OpenFlags(flags)
}

type getxattrIn struct {
	getxattrInCommon
}

type setxattrIn struct {
	setxattrInCommon
}
