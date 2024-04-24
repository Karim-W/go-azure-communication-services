package logger

import "time"

func String(key, value string) LoggerField {
	return LoggerField{
		Key:   key,
		Value: value,
		Type:  "string",
	}
}

func Int64(key string, value int64) LoggerField {
	return LoggerField{
		Key:   key,
		Value: value,
		Type:  "int64",
	}
}

func Int(key string, value int) LoggerField {
	return LoggerField{
		Key:   key,
		Value: value,
		Type:  "int",
	}
}

func Uint64(key string, v uint64) LoggerField {
	return LoggerField{
		Key:   key,
		Value: v,
		Type:  "uint64",
	}
}

func Float64(key string, v float64) LoggerField {
	return LoggerField{
		Key:   key,
		Value: v,
		Type:  "float64",
	}
}

func Bool(key string, v bool) LoggerField {
	return LoggerField{
		Key:   key,
		Value: v,
		Type:  "bool",
	}
}

func Time(key string, v time.Time) LoggerField {
	return LoggerField{
		Key:   key,
		Value: v,
		Type:  "time",
	}
}

func Duration(key string, v time.Duration) LoggerField {
	return LoggerField{
		Key:   key,
		Value: v,
		Type:  "duration",
	}
}
