package logger

import (
	"errors"
	"strings"
)

const (
	Error Level = iota + 1
	Warn
	Info
	Debug
)

var ErrInvalidLogLevel = errors.New("unrecognized log level")

type Level int

var levels = map[Level]string{
	Error: "error",
	Warn:  "warn",
	Info:  "info",
	Debug: "debug",
}

func (lvl Level) String() string {
	return levels[lvl]
}

func (lvl Level) isAllowed(logLevel Level) bool {
	return lvl <= logLevel
}

func (lvl *Level) UnmarshalText(text string) error {
	switch string(strings.ToLower(text)) {
	case "debug":
		*lvl = Debug
	case "info":
		*lvl = Info
	case "warn":
		*lvl = Warn
	case "error":
		*lvl = Error
	default:
		return ErrInvalidLogLevel
	}
	return nil
}
