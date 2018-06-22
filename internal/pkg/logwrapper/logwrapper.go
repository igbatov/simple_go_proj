package logwrapper

import "simple_go_proj/internal/pkg/log"

type CanLog interface {
	Log()
}

type LogWrapper struct {
	L CanLog
}

func New() LogWrapper {
	logger := log.New()
	return LogWrapper{L: logger}
}
