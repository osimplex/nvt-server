package worker

import (
	"nvt-server/src/model"
	"nvt-server/src/tracer"
)

type Worker interface {
	GetFrame(obj *model.Frame, tracer tracer.Tracer) error
	GetFrameCommand(obj *model.FrameCommand, tracer tracer.Tracer) error
	ProcedureCall(obj *model.Session, tracer tracer.Tracer) error

	End() error
}
