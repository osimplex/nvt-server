package handler

import (
	"context"
	"errors"
	"io"
	"nvt-server/src/common"
	"nvt-server/src/model"
	"nvt-server/src/tui"
	"nvt-server/src/vt100"
	"strings"

	"github.com/reiver/go-oi"
)

var (
	errEndOfInteraction = errors.New("end of interaction")
	errContinueLoop     = errors.New("continue loop")
)

type TuiHandler struct {
	Config   *common.Config
	ctxValue *ContextValue
}

func (handler TuiHandler) HandleConn(ctx context.Context, w io.Writer, r io.Reader) {
	handler.ctxValue = ctx.Value(common.CtxKey).(*ContextValue)

	var (
		inputBuffer []byte
		printInput  = false

		conn    = handler.ctxValue.Conn
		session = handler.ctxValue.Session
		tracer  = handler.ctxValue.Tracer
		config  = handler.Config.HandlerOpts[handler.Config.Handler]

		frameBuffer  = &model.Frame{}
		frameCommand = &model.FrameCommand{}
	)

	session.ClientAddr = strings.Split(conn.RemoteAddr().String(), ":")[0]

	oi.LongWriteString(w, tui.MessageOfTheDay())

handlerLoop:
	for {
		sym, err := vt100.KeyEval(r)

		if err != nil {
			switch err {
			case io.EOF:
				tracer.Tracef("request of end of communication, conn %q", conn.RemoteAddr())
				break handlerLoop
			case vt100.ErrChildNotFound:
				tracer.Debugf("error in %q communication at symbol processing: %q", conn.RemoteAddr(), err)
			case vt100.ErrNoChildren:
				tracer.Debugf("error in %q communication at symbol processing: %q", conn.RemoteAddr(), err)
			default:
				tracer.Errorf("ending communication with %q, unknown error: %q", conn.RemoteAddr(), err)
				break handlerLoop
			}
		}

		switch sym.Type {
		case vt100.SymCommand:
			frameCommand.Command = config.Opts[sym.Command]
			session.Input = string(inputBuffer)
			inputBuffer = []byte{}
		case vt100.SymDigit:
			inputBuffer = append(inputBuffer, sym.Symbol)
			if printInput {
				oi.LongWriteByte(w, sym.Symbol)
			}
			continue handlerLoop
		default:
			continue handlerLoop
		}

		tracer.Tracef(
			"receiving command %q, user id %q on %q",
			frameCommand.Command,
			session.Register.Rv01,
			conn.RemoteAddr(),
		)

		switch handler.commandEval(frameCommand) {
		case errEndOfInteraction:
			break handlerLoop
		case errContinueLoop:
			continue handlerLoop
		}

		switch handler.frameRecv(frameBuffer) {
		case errContinueLoop:
			continue handlerLoop
		}

		cursorPositionSequence, err := tui.DetermineCursorPosition(*frameBuffer)
		printInput = err == nil
		newFrame := tui.BuildFrameString(*frameBuffer, session.FrameLines)

		oi.LongWriteString(w, tui.ClearScreenSequence)
		oi.LongWriteString(w, newFrame)
		oi.LongWriteString(w, cursorPositionSequence)
	}
}

func (handler TuiHandler) commandEval(frameCommand *model.FrameCommand) error {
	var (
		conn    = handler.ctxValue.Conn
		session = handler.ctxValue.Session
		tracer  = handler.ctxValue.Tracer
		worker  = handler.ctxValue.Worker
		config  = handler.Config.HandlerOpts[handler.Config.Handler]

		firstFrame      = config.Opts["FirstFrame"]
		exitTransaction = config.Opts["ExitTransaction"]
	)

	if session.FrameName == "" {
		frameCommand.FrameName = firstFrame
		session.FrameName = firstFrame

		return nil
	}

	frameCommand.FrameName = session.FrameName
	if err := worker.GetFrameCommand(frameCommand, tracer); err != nil {
		return errContinueLoop
	}

	switch {
	case (frameCommand.FrameRef != "") == (frameCommand.Transaction != ""):
		tracer.Warnf(
			"%q does not have valid actions for command %q, or is ambiguous",
			frameCommand.FrameName,
			frameCommand.Command,
		)
		return errContinueLoop
	case frameCommand.Transaction == exitTransaction:
		tracer.Tracef(
			"request of end of interaction, user id %q on %q",
			session.Register.Rv01,
			conn.RemoteAddr(),
		)
		return errEndOfInteraction
	}

	session.PreviousFrameName = session.FrameName
	if frameCommand.Transaction != "" {
		session.Transaction = frameCommand.Transaction
		tracer.Tracef(
			"running %q (with input %q), user id %q on %q",
			session.Transaction,
			session.Input,
			session.Register.Rv01,
			conn.RemoteAddr(),
		)
		if err := worker.ProcedureCall(session, tracer); err != nil {
			return errContinueLoop
		}
	} else {
		session.FrameName = frameCommand.FrameRef
	}

	return nil
}

func (handler TuiHandler) frameRecv(frameBuffer *model.Frame) error {
	var (
		conn    = handler.ctxValue.Conn
		session = handler.ctxValue.Session
		tracer  = handler.ctxValue.Tracer
		worker  = handler.ctxValue.Worker
	)

	frameBuffer.FrameName = session.FrameName
	tracer.Tracef(
		"recovering frame %q (previous frame %q), user id %q on %q",
		frameBuffer.FrameName,
		session.PreviousFrameName,
		session.Register.Rv01,
		conn.RemoteAddr(),
	)
	if err := worker.GetFrame(frameBuffer, tracer); err != nil {
		return errContinueLoop
	}

	return nil
}
