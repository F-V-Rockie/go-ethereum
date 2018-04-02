package fuse

// Unmount tries to unmount the filesystem mounted at dir.
func Unmount(dir string) error { log.DebugLog()
	return unmount(dir)
}
