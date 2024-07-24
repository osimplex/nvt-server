package tracer

import (
	"log"
	"nvt-server/src/common"
)

type TraceMode int

const (
	Debug TraceMode = iota
	Trace
	Warn
	Error
)

type LocalTracer struct {
	Config *common.Config
}

func (tracer *LocalTracer) Debug(v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Debug {
		log.Println(v...)
	}
}
func (tracer *LocalTracer) Debugf(format string, v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Debug {
		log.Printf(format+"\n", v...)
	}
}

func (tracer *LocalTracer) Trace(v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Trace {
		log.Println(v...)
	}
}
func (tracer *LocalTracer) Tracef(format string, v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Trace {
		log.Printf(format+"\n", v...)
	}
}

func (tracer *LocalTracer) Warn(v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Warn {
		log.Println(v...)
	}
}
func (tracer *LocalTracer) Warnf(format string, v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Warn {
		log.Printf(format+"\n", v...)
	}
}

func (tracer *LocalTracer) Error(v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Error {
		log.Println(v...)
	}
}
func (tracer *LocalTracer) Errorf(format string, v ...any) {
	if TraceMode(tracer.Config.LogMode) <= Error {
		log.Printf(format+"\n", v...)
	}
}
