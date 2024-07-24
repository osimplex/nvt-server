package server

import (
	"context"
	"errors"
	"net"
	"nvt-server/src/common"
	"nvt-server/src/handler"
	"nvt-server/src/model"
	"nvt-server/src/tracer"
	"nvt-server/src/worker"
	"sync"

	"github.com/ziutek/telnet"
)

type Telnet struct {
	Addr    string
	Handler handler.Handler
	Worker  worker.Worker
	Tracer  tracer.Tracer

	stop    chan struct{}
	nospawn chan struct{}
	stoped  chan struct{}
}

func (server *Telnet) End() {
	close(server.stop)
	<-server.stoped
}

func (server *Telnet) ListenAndServe() error {
	var (
		handlersWaitGroup sync.WaitGroup
		serverTracer      = server.Tracer
	)

	ln, err := net.Listen("tcp", server.Addr)
	if err != nil {
		return err
	}
	serverTracer.Tracef("listening at %q", ln.Addr())

	server.stop = make(chan struct{})
	server.nospawn = make(chan struct{})
	server.stoped = make(chan struct{})

	go func() {
		<-server.stop
		serverTracer.Trace("ending telnet server...")
		ln.Close()
		handlersWaitGroup.Wait()
		<-server.nospawn
		if server.Worker != nil {
			serverTracer.Trace("ending server's worker...")
			server.Worker.End()
		}
		close(server.stoped)
	}()

	go func() {
	connLoop:
		for {
			conn, err := ln.Accept()
			switch {
			case errors.Is(err, net.ErrClosed):
				serverTracer.Errorf("conection closed: %q", err)
				close(server.nospawn)
				break connLoop
			case err != nil:
				serverTracer.Errorf("unknown error receiving connection: %q", err)
				break connLoop
			}
			serverTracer.Tracef("received new connection from %q.", conn.RemoteAddr())

			handlersWaitGroup.Add(1)
			go func() {
				server.handle(conn)
				handlersWaitGroup.Done()
			}()
		}
	}()

	return nil
}

func (server *Telnet) handle(conn net.Conn) {
	defer conn.Close()

	defer func() {
		if r := recover(); nil != r {
			server.Tracer.Warnf("recovered from: (%T) %v", r, r)
		}
	}()

	ctxValue := &handler.ContextValue{
		Conn:    conn,
		Session: &model.Session{},
		Tracer:  server.Tracer,
		Worker:  server.Worker,
	}

	telnetConn, err := telnet.NewConn(conn)
	if err != nil {
		server.Tracer.Errorf("telnet conn establishing error: %q", err)
	}

	ctx := context.WithValue(context.Background(), common.CtxKey, ctxValue)

	server.Handler.HandleConn(ctx, telnetConn, telnetConn)
}
