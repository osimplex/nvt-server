package sql

import (
	"nvt-server/src/model"

	"gopkg.in/guregu/null.v4"
)

type FrameCommand struct {
	FrameName   null.String
	Command     null.String
	Transaction null.String
	FrameRef    null.String
}

func (FrameCommand) GetDto() FrameCommand {
	return FrameCommand{}
}

func (dtoFrameCommand *FrameCommand) GetDtoPtr() *FrameCommand {
	return dtoFrameCommand
}

func (dtoFrameCommand *FrameCommand) SetDto(src *model.FrameCommand) {
	dtoFrameCommand.FrameName = null.StringFromPtr(&src.FrameName)
	dtoFrameCommand.Command = null.StringFromPtr(&src.Command)
	dtoFrameCommand.Transaction = null.StringFromPtr(&src.Transaction)
	dtoFrameCommand.FrameRef = null.StringFromPtr(&src.FrameRef)
}

func (dtoFrameCommand *FrameCommand) SetObject(dest *model.FrameCommand) {
	dest.FrameName = dtoFrameCommand.FrameName.ValueOrZero()
	dest.Command = dtoFrameCommand.Command.ValueOrZero()
	dest.Transaction = dtoFrameCommand.Transaction.ValueOrZero()
	dest.FrameRef = dtoFrameCommand.FrameRef.ValueOrZero()
}
