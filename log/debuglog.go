package log

import (
	"runtime"
	"strings"
	"log"
)

func DebugLog() {
	pc, file, line, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")[2]
	log.Printf("%s/%s() line %d", file, f, line)
}
