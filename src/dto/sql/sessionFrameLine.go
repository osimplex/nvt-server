package sql

import (
	"nvt-server/src/model"

	"gopkg.in/guregu/null.v4"
)

type SessionFrameLine struct {
	Rl1 null.String
	Rl2 null.String
	Rl3 null.String
	Rl4 null.String
	Rl5 null.String
	Rl6 null.String
	Rl7 null.String
	Rl8 null.String
}

func (dtoSessionFrameLine *SessionFrameLine) SetDto(src *model.SessionFrameLine) {
	dtoSessionFrameLine.Rl1 = null.StringFromPtr(&src.Rl1)
	dtoSessionFrameLine.Rl2 = null.StringFromPtr(&src.Rl2)
	dtoSessionFrameLine.Rl3 = null.StringFromPtr(&src.Rl3)
	dtoSessionFrameLine.Rl4 = null.StringFromPtr(&src.Rl4)
	dtoSessionFrameLine.Rl5 = null.StringFromPtr(&src.Rl5)
	dtoSessionFrameLine.Rl6 = null.StringFromPtr(&src.Rl6)
	dtoSessionFrameLine.Rl7 = null.StringFromPtr(&src.Rl7)
	dtoSessionFrameLine.Rl8 = null.StringFromPtr(&src.Rl8)
}

func (dtoSessionFrameLine *SessionFrameLine) SetObject(dest *model.SessionFrameLine) {
	dest.Rl1 = dtoSessionFrameLine.Rl1.ValueOrZero()
	dest.Rl2 = dtoSessionFrameLine.Rl2.ValueOrZero()
	dest.Rl3 = dtoSessionFrameLine.Rl3.ValueOrZero()
	dest.Rl4 = dtoSessionFrameLine.Rl4.ValueOrZero()
	dest.Rl5 = dtoSessionFrameLine.Rl5.ValueOrZero()
	dest.Rl6 = dtoSessionFrameLine.Rl6.ValueOrZero()
	dest.Rl7 = dtoSessionFrameLine.Rl7.ValueOrZero()
	dest.Rl8 = dtoSessionFrameLine.Rl8.ValueOrZero()
}
