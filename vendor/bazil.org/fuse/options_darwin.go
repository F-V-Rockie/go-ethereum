package fuse

func localVolume(conf *mountConfig) error { log.DebugLog()
	conf.options["local"] = ""
	return nil
}

func volumeName(name string) MountOption { log.DebugLog()
	return func(conf *mountConfig) error {
		conf.options["volname"] = name
		return nil
	}
}

func daemonTimeout(name string) MountOption { log.DebugLog()
	return func(conf *mountConfig) error {
		conf.options["daemon_timeout"] = name
		return nil
	}
}

func noAppleXattr(conf *mountConfig) error { log.DebugLog()
	conf.options["noapplexattr"] = ""
	return nil
}

func noAppleDouble(conf *mountConfig) error { log.DebugLog()
	conf.options["noappledouble"] = ""
	return nil
}

func exclCreate(conf *mountConfig) error { log.DebugLog()
	conf.options["excl_create"] = ""
	return nil
}
