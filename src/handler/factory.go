package handler

import (
	"errors"
	"nvt-server/src/common"
)

var (
	ErrUnknownHandler = errors.New("unknown handler requested")
)

func Factory(config *common.Config) (Handler, error) {
	switch config.Handler {
	case "tui":
		return &TuiHandler{
			Config: config,
		}, nil
	case "example":
		return &ExampleHandler{}, nil
	}

	return nil, ErrUnknownHandler
}
