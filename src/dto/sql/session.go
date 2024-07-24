package sql

import (
	"nvt-server/src/model"

	"gopkg.in/guregu/null.v4"
)

type Session struct {
	ClientAddr        null.String
	Transaction       null.String
	Input             null.String
	SearchType        null.String
	Beep              null.String
	FrameName         null.String
	PreviousFrameName null.String

	Register   SessionRegister
	FrameLines SessionFrameLine
}

func (Session) GetDto() Session {
	return Session{}
}

func (dtoSession *Session) GetDtoPtr() *Session {
	return dtoSession
}

func (dtoSession *Session) SetDto(src *model.Session) {
	dtoSession.ClientAddr = null.StringFromPtr(&src.ClientAddr)
	dtoSession.Transaction = null.StringFromPtr(&src.Transaction)
	dtoSession.Input = null.StringFromPtr(&src.Input)
	dtoSession.SearchType = null.StringFromPtr(&src.SearchType)
	dtoSession.Beep = null.StringFromPtr(&src.Beep)
	dtoSession.FrameName = null.StringFromPtr(&src.FrameName)
	dtoSession.PreviousFrameName = null.StringFromPtr(&src.PreviousFrameName)

	dtoSession.Register.SetDto(&src.Register)
	dtoSession.FrameLines.SetDto(&src.FrameLines)
}

func (dtoSession *Session) SetObject(dest *model.Session) {
	dest.ClientAddr = dtoSession.ClientAddr.ValueOrZero()
	dest.Transaction = dtoSession.Transaction.ValueOrZero()
	dest.Input = dtoSession.Input.ValueOrZero()
	dest.SearchType = dtoSession.SearchType.ValueOrZero()
	dest.Beep = dtoSession.Beep.ValueOrZero()
	dest.FrameName = dtoSession.FrameName.ValueOrZero()
	dest.PreviousFrameName = dtoSession.PreviousFrameName.ValueOrZero()

	dtoSession.Register.SetObject(&dest.Register)
	dtoSession.FrameLines.SetObject(&dest.FrameLines)
}
