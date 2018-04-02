package otto

import (
	"strconv"
)

func (runtime *_runtime) newBooleanObject(value Value) *_object { log.DebugLog()
	return runtime.newPrimitiveObject("Boolean", toValue_bool(value.bool()))
}

func booleanToString(value bool) string { log.DebugLog()
	return strconv.FormatBool(value)
}
