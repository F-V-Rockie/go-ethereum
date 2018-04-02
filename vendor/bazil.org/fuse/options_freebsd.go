package fuse

func localVolume(conf *mountConfig) error { log.DebugLog()
	return nil
}

func volumeName(name string) MountOption { log.DebugLog()
	return dummyOption
}

func daemonTimeout(name string) MountOption { log.DebugLog()
	return func(conf *mountConfig) error {
		conf.options["timeout"] = name
		return nil
	}
}

func noAppleXattr(conf *mountConfig) error { log.DebugLog()
	return nil
}

func noAppleDouble(conf *mountConfig) error { log.DebugLog()
	return nil
}

func exclCreate(conf *mountConfig) error { log.DebugLog()
	return nil
}
