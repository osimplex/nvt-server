package oracle

import (
	"database/sql"
	"nvt-server/src/common"
	dto "nvt-server/src/dto/sql"
)

type Oracle struct {
	DtoFrame        *dto.Frame
	DtoFrameCommand *dto.FrameCommand
	DtoSession      *dto.Session
	Config          *common.Config
	DB              *sql.DB
}

func (oracle *Oracle) End() error {
	return oracle.DB.Close()
}
