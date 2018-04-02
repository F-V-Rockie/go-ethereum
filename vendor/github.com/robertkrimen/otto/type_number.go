package otto

func (runtime *_runtime) newNumberObject(value Value) *_object { log.DebugLog()
	return runtime.newPrimitiveObject("Number", value.numberValue())
}
