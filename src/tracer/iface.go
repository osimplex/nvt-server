package tracer

type Tracer interface {
	Debug(v ...any)
	Debugf(format string, v ...any)

	Trace(v ...any)
	Tracef(format string, v ...any)

	Warn(v ...any)
	Warnf(format string, v ...any)

	Error(v ...any)
	Errorf(format string, v ...any)
}
