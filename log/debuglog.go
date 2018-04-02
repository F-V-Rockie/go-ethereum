package log

import (
	"runtime"
	"strings"
	"log"
)

func DebugLog() { log.DebugLog()
	pc, file1, line, _ := runtime.Caller(1)
	f := strings.Split(runtime.FuncForPC(pc).Name(), ".")[1]

	log.Println(file1, "/", f, "()", "line", line)
}
