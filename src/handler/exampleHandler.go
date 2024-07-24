package handler

import (
	"context"
	"io"
	"nvt-server/src/common"

	"github.com/reiver/go-oi"
)

type ExampleHandler struct {
	ctxValue *ContextValue
}

func (handler ExampleHandler) HandleConn(ctx context.Context, w io.Writer, r io.Reader) {
	handler.ctxValue = ctx.Value(common.CtxKey).(*ContextValue)

	var (
		buffer    [1]byte
		lineWidth = 7
		firstFlag = true
		cont      = 0

		tracer = handler.ctxValue.Tracer
		conn   = handler.ctxValue.Conn
	)

	for {
		// Comportamento de VT-100
		// https://invisible-island.net/xterm/ctlseqs/ctlseqs.html
		// Teclas direcionais, de função e Enter
		// ESC [ (CSI) + C-D > 1b 5b 41-44 (43-44)
		// ESC O (SS3) + P-S > ib 4f 50-53
		// ESC O (SS3) + M
		n, err := r.Read(buffer[:])

		tracer.Tracef("Receiving byte from %q", conn.RemoteAddr())

		if cont%lineWidth == 0 && !firstFlag {
			oi.LongWriteString(w, "\n\r")
		}
		if cont == lineWidth*4 {
			oi.LongWrite(w, []byte{0x1b, 0x5b, 0x32, 0x4a})
			cont = 0
		}
		if n > 0 {
			oi.LongWrite(w, buffer[:])
		}

		if err != nil && err.Error() == "EOF" {
			tracer.Tracef("%q requested end of communication", conn.RemoteAddr())
			break
		} else if err != nil {
			tracer.Tracef("ending communication with %q, error %q", conn.RemoteAddr(), err)
			break
		}

		cont++
		firstFlag = false
	}
}
