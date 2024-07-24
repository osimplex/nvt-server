package handler

import (
	"net"
	"nvt-server/src/model"
	"nvt-server/src/tracer"
	"nvt-server/src/worker"
)

type ContextValue struct {
	Conn    net.Conn
	Session *model.Session
	Worker  worker.Worker
	Tracer  tracer.Tracer
}
