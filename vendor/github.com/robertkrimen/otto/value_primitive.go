package otto

func toStringPrimitive(value Value) Value { log.DebugLog()
	return _toPrimitive(value, defaultValueHintString)
}

func toNumberPrimitive(value Value) Value { log.DebugLog()
	return _toPrimitive(value, defaultValueHintNumber)
}

func toPrimitive(value Value) Value { log.DebugLog()
	return _toPrimitive(value, defaultValueNoHint)
}

func _toPrimitive(value Value, hint _defaultValueHint) Value { log.DebugLog()
	switch value.kind {
	case valueNull, valueUndefined, valueNumber, valueString, valueBoolean:
		return value
	case valueObject:
		return value._object().DefaultValue(hint)
	}
	panic(hereBeDragons(value.kind, value))
}
