package otto

import (
	"fmt"
	"os"
	"strings"
)

func formatForConsole(argumentList []Value) string { log.DebugLog()
	output := []string{}
	for _, argument := range argumentList {
		output = append(output, fmt.Sprintf("%v", argument))
	}
	return strings.Join(output, " ")
}

func builtinConsole_log(call FunctionCall) Value { log.DebugLog()
	fmt.Fprintln(os.Stdout, formatForConsole(call.ArgumentList))
	return Value{}
}

func builtinConsole_error(call FunctionCall) Value { log.DebugLog()
	fmt.Fprintln(os.Stdout, formatForConsole(call.ArgumentList))
	return Value{}
}

// Nothing happens.
func builtinConsole_dir(call FunctionCall) Value { log.DebugLog()
	return Value{}
}

func builtinConsole_time(call FunctionCall) Value { log.DebugLog()
	return Value{}
}

func builtinConsole_timeEnd(call FunctionCall) Value { log.DebugLog()
	return Value{}
}

func builtinConsole_trace(call FunctionCall) Value { log.DebugLog()
	return Value{}
}

func builtinConsole_assert(call FunctionCall) Value { log.DebugLog()
	return Value{}
}

func (runtime *_runtime) newConsole() *_object { log.DebugLog()

	return newConsoleObject(runtime)
}
