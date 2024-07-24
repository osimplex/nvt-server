package server

import (
	"errors"
	"nvt-server/src/common"
	"nvt-server/src/handler"
	"nvt-server/src/tracer"
	"nvt-server/src/worker"
)

var (
	ErrUnknownServer = errors.New("unknown server requested")
)

func Factory(config *common.Config) (Server, error) {
	var (
		serverWorker  worker.Worker
		serverHandler handler.Handler
		serverTracer  tracer.Tracer
		err           error
	)

	serverTracer = &tracer.LocalTracer{
		Config: config,
	}

	switch config.Server {
	case "telnet":
		if config.Worker != "" {
			serverWorker, err = worker.Factory(config)
			if err != nil {
				return nil, err
			}
		}

		serverHandler, err = handler.Factory(config)
		if err != nil {
			return nil, err
		}

		return &Telnet{
			Addr:    config.ServerOpts[config.Server].Addr,
			Handler: serverHandler,
			Worker:  serverWorker,
			Tracer:  serverTracer,
		}, nil
	}

	return nil, ErrUnknownServer
}
