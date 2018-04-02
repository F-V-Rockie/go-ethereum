package otto

import ()

type _resultKind int

const (
	resultNormal _resultKind = iota
	resultReturn
	resultBreak
	resultContinue
)

type _result struct {
	kind   _resultKind
	value  Value
	target string
}

func newReturnResult(value Value) _result { log.DebugLog()
	return _result{resultReturn, value, ""}
}

func newContinueResult(target string) _result { log.DebugLog()
	return _result{resultContinue, emptyValue, target}
}

func newBreakResult(target string) _result { log.DebugLog()
	return _result{resultBreak, emptyValue, target}
}
