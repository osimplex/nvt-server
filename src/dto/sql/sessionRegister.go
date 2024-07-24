package sql

import (
	"nvt-server/src/model"

	"gopkg.in/guregu/null.v4"
)

type SessionRegister struct {
	Rv01 null.String
	Rv02 null.String
	Rv03 null.String
	Rv04 null.String
	Rv05 null.String
	Rv06 null.String
	Rv07 null.String
	Rv08 null.String
	Rv09 null.String
	Rv10 null.String
	Rv11 null.String
	Rv12 null.String
	Rv13 null.String
	Rv14 null.String
	Rv15 null.String
	Rv16 null.String
	Rv17 null.String
	Rv18 null.String
	Rv19 null.String
	Rv20 null.String
	Rv21 null.String
	Rv22 null.String
	Rv23 null.String
	Rv24 null.String
	Rv25 null.String
	Rv26 null.String
	Rv27 null.String
	Rv28 null.String
	Rv29 null.String
	Rv30 null.String
	Rv31 null.String
	Rv32 null.String
}

func (dtoSessionRegister *SessionRegister) SetDto(src *model.SessionRegister) {
	dtoSessionRegister.Rv01 = null.StringFromPtr(&src.Rv01)
	dtoSessionRegister.Rv02 = null.StringFromPtr(&src.Rv02)
	dtoSessionRegister.Rv03 = null.StringFromPtr(&src.Rv03)
	dtoSessionRegister.Rv04 = null.StringFromPtr(&src.Rv04)
	dtoSessionRegister.Rv05 = null.StringFromPtr(&src.Rv05)
	dtoSessionRegister.Rv06 = null.StringFromPtr(&src.Rv06)
	dtoSessionRegister.Rv07 = null.StringFromPtr(&src.Rv07)
	dtoSessionRegister.Rv08 = null.StringFromPtr(&src.Rv08)
	dtoSessionRegister.Rv09 = null.StringFromPtr(&src.Rv09)
	dtoSessionRegister.Rv10 = null.StringFromPtr(&src.Rv10)
	dtoSessionRegister.Rv11 = null.StringFromPtr(&src.Rv11)
	dtoSessionRegister.Rv12 = null.StringFromPtr(&src.Rv12)
	dtoSessionRegister.Rv13 = null.StringFromPtr(&src.Rv13)
	dtoSessionRegister.Rv14 = null.StringFromPtr(&src.Rv14)
	dtoSessionRegister.Rv15 = null.StringFromPtr(&src.Rv15)
	dtoSessionRegister.Rv16 = null.StringFromPtr(&src.Rv16)
	dtoSessionRegister.Rv17 = null.StringFromPtr(&src.Rv17)
	dtoSessionRegister.Rv18 = null.StringFromPtr(&src.Rv18)
	dtoSessionRegister.Rv19 = null.StringFromPtr(&src.Rv19)
	dtoSessionRegister.Rv20 = null.StringFromPtr(&src.Rv20)
	dtoSessionRegister.Rv21 = null.StringFromPtr(&src.Rv21)
	dtoSessionRegister.Rv22 = null.StringFromPtr(&src.Rv22)
	dtoSessionRegister.Rv23 = null.StringFromPtr(&src.Rv23)
	dtoSessionRegister.Rv24 = null.StringFromPtr(&src.Rv24)
	dtoSessionRegister.Rv25 = null.StringFromPtr(&src.Rv25)
	dtoSessionRegister.Rv26 = null.StringFromPtr(&src.Rv26)
	dtoSessionRegister.Rv27 = null.StringFromPtr(&src.Rv27)
	dtoSessionRegister.Rv28 = null.StringFromPtr(&src.Rv28)
	dtoSessionRegister.Rv29 = null.StringFromPtr(&src.Rv29)
	dtoSessionRegister.Rv30 = null.StringFromPtr(&src.Rv30)
	dtoSessionRegister.Rv31 = null.StringFromPtr(&src.Rv31)
	dtoSessionRegister.Rv32 = null.StringFromPtr(&src.Rv32)
}

func (dtoSessionRegister *SessionRegister) SetObject(dest *model.SessionRegister) {
	dest.Rv01 = dtoSessionRegister.Rv01.ValueOrZero()
	dest.Rv02 = dtoSessionRegister.Rv02.ValueOrZero()
	dest.Rv03 = dtoSessionRegister.Rv03.ValueOrZero()
	dest.Rv04 = dtoSessionRegister.Rv04.ValueOrZero()
	dest.Rv05 = dtoSessionRegister.Rv05.ValueOrZero()
	dest.Rv06 = dtoSessionRegister.Rv06.ValueOrZero()
	dest.Rv07 = dtoSessionRegister.Rv07.ValueOrZero()
	dest.Rv08 = dtoSessionRegister.Rv08.ValueOrZero()
	dest.Rv09 = dtoSessionRegister.Rv09.ValueOrZero()
	dest.Rv10 = dtoSessionRegister.Rv10.ValueOrZero()
	dest.Rv11 = dtoSessionRegister.Rv11.ValueOrZero()
	dest.Rv12 = dtoSessionRegister.Rv12.ValueOrZero()
	dest.Rv13 = dtoSessionRegister.Rv13.ValueOrZero()
	dest.Rv14 = dtoSessionRegister.Rv14.ValueOrZero()
	dest.Rv15 = dtoSessionRegister.Rv15.ValueOrZero()
	dest.Rv16 = dtoSessionRegister.Rv16.ValueOrZero()
	dest.Rv17 = dtoSessionRegister.Rv17.ValueOrZero()
	dest.Rv18 = dtoSessionRegister.Rv18.ValueOrZero()
	dest.Rv19 = dtoSessionRegister.Rv19.ValueOrZero()
	dest.Rv20 = dtoSessionRegister.Rv20.ValueOrZero()
	dest.Rv21 = dtoSessionRegister.Rv21.ValueOrZero()
	dest.Rv22 = dtoSessionRegister.Rv22.ValueOrZero()
	dest.Rv23 = dtoSessionRegister.Rv23.ValueOrZero()
	dest.Rv24 = dtoSessionRegister.Rv24.ValueOrZero()
	dest.Rv25 = dtoSessionRegister.Rv25.ValueOrZero()
	dest.Rv26 = dtoSessionRegister.Rv26.ValueOrZero()
	dest.Rv27 = dtoSessionRegister.Rv27.ValueOrZero()
	dest.Rv28 = dtoSessionRegister.Rv28.ValueOrZero()
	dest.Rv29 = dtoSessionRegister.Rv29.ValueOrZero()
	dest.Rv30 = dtoSessionRegister.Rv30.ValueOrZero()
	dest.Rv31 = dtoSessionRegister.Rv31.ValueOrZero()
	dest.Rv32 = dtoSessionRegister.Rv32.ValueOrZero()
}
