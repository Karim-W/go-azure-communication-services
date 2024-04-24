package logger

import (
	"log/slog"
	"time"
)

type LoggerField struct {
	Key   string
	Value any
	Type  string
}

type Logger interface {
	Info(
		msg string,
		fields ...LoggerField,
	)
	Error(
		msg string,
		fields ...LoggerField,
	)
}

type _noopLogger struct{}

func (_ *_noopLogger) Error(msg string, fields ...LoggerField) {
}

func (_ *_noopLogger) Info(msg string, fields ...LoggerField) {
}

func Noop() Logger {
	return &_noopLogger{}
}

type defaultLogger struct{}

var _ Logger = &defaultLogger{}

func Default() Logger {
	return &defaultLogger{}
}

func (l *defaultLogger) Info(
	msg string,
	fields ...LoggerField,
) {
	f := make([]any, len(fields))

	for i, field := range fields {
		switch field.Type {
		case "string":
			f[i] = slog.String(field.Key, field.Value.(string))
		case "int64":
			f[i] = slog.Int64(field.Key, field.Value.(int64))
		case "int":
			f[i] = slog.Int(field.Key, field.Value.(int))
		case "uint64":
			f[i] = slog.Uint64(field.Key, field.Value.(uint64))
		case "float64":
			f[i] = slog.Float64(field.Key, field.Value.(float64))
		case "bool":
			f[i] = slog.Bool(field.Key, field.Value.(bool))
		case "time":
			f[i] = slog.Time(field.Key, field.Value.(time.Time))
		case "duration":
			f[i] = slog.Duration(field.Key, field.Value.(time.Duration))
		}
	}

	slog.Info(msg, f...)
}

func (l *defaultLogger) Error(
	msg string,
	fields ...LoggerField,
) {
	f := make([]any, len(fields))

	for i, field := range fields {
		switch field.Type {
		case "string":
			f[i] = slog.String(field.Key, field.Value.(string))
		case "int64":
			f[i] = slog.Int64(field.Key, field.Value.(int64))
		case "int":
			f[i] = slog.Int(field.Key, field.Value.(int))
		case "uint64":
			f[i] = slog.Uint64(field.Key, field.Value.(uint64))
		case "float64":
			f[i] = slog.Float64(field.Key, field.Value.(float64))
		case "bool":
			f[i] = slog.Bool(field.Key, field.Value.(bool))
		case "time":
			f[i] = slog.Time(field.Key, field.Value.(time.Time))
		case "duration":
			f[i] = slog.Duration(field.Key, field.Value.(time.Duration))
		}
	}

	slog.Error(msg, f...)
}
