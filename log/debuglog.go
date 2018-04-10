package log

import (
	"runtime"
	"strings"
	"log"
)

func DebugLog() {
	pc, file, line, _ := runtime.Caller(1)
	f := getName(pc)
	log.Printf("%s/%s() line %d", file, f, line)
}

func getName(_pc uintptr) string {
	longname := strings.Split(runtime.FuncForPC(_pc).Name(), "/")
	fullname := longname[len(longname)-1]
	funcname := strings.SplitN(fullname, ".", 2)
	return funcname[1]
}
