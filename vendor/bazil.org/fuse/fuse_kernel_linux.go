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
	// Ignored on Linux.
}

func (a *attr) SetFlags(f uint32) { log.DebugLog()
	// Ignored on Linux.
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
	// on amd64, the 32-bit O_LARGEFILE flag is always seen;
	// on i386, the flag probably depends on the app
	// requesting, but in any case should be utterly
	// uninteresting to us here; our kernel protocol messages
	// are not directly related to the client app's kernel
	// API/ABI
	flags &^= 0x8000

	return OpenFlags(flags)
}

type getxattrIn struct {
	getxattrInCommon
}

type setxattrIn struct {
	setxattrInCommon
}
