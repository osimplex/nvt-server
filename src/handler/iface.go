package handler

import (
	"context"
	"io"
)

type Handler interface {
	HandleConn(ctx context.Context, w io.Writer, r io.Reader)
}
