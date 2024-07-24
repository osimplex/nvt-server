package sql

import (
	"nvt-server/src/model"

	"gopkg.in/guregu/null.v4"
)

type Frame struct {
	FrameName  null.String
	Class      null.String
	DeviceType null.String
	InputFrame null.String

	Ln1 null.String
	Ln2 null.String
	Ln3 null.String
	Ln4 null.String
	Ln5 null.String
	Ln6 null.String
	Ln7 null.String
	Ln8 null.String
}

func (Frame) GetDto() Frame {
	return Frame{}
}

func (frame *Frame) GetDtoPtr() *Frame {
	return frame
}

func (dtoFrame *Frame) SetDto(src *model.Frame) {
	dtoFrame.FrameName = null.StringFromPtr(&src.FrameName)
	dtoFrame.Class = null.StringFromPtr(&src.Class)
	dtoFrame.DeviceType = null.StringFromPtr(&src.DeviceType)
	dtoFrame.InputFrame = null.StringFromPtr(&src.InputFrame)

	dtoFrame.Ln1 = null.StringFromPtr(&src.Ln1)
	dtoFrame.Ln2 = null.StringFromPtr(&src.Ln2)
	dtoFrame.Ln3 = null.StringFromPtr(&src.Ln3)
	dtoFrame.Ln4 = null.StringFromPtr(&src.Ln4)
	dtoFrame.Ln5 = null.StringFromPtr(&src.Ln5)
	dtoFrame.Ln6 = null.StringFromPtr(&src.Ln6)
	dtoFrame.Ln7 = null.StringFromPtr(&src.Ln7)
	dtoFrame.Ln8 = null.StringFromPtr(&src.Ln8)
}

func (dtoFrame *Frame) SetObject(dest *model.Frame) {
	dest.FrameName = dtoFrame.FrameName.ValueOrZero()
	dest.Class = dtoFrame.Class.ValueOrZero()
	dest.DeviceType = dtoFrame.DeviceType.ValueOrZero()
	dest.InputFrame = dtoFrame.InputFrame.ValueOrZero()

	dest.Ln1 = dtoFrame.Ln1.ValueOrZero()
	dest.Ln2 = dtoFrame.Ln2.ValueOrZero()
	dest.Ln3 = dtoFrame.Ln3.ValueOrZero()
	dest.Ln4 = dtoFrame.Ln4.ValueOrZero()
	dest.Ln5 = dtoFrame.Ln5.ValueOrZero()
	dest.Ln6 = dtoFrame.Ln6.ValueOrZero()
	dest.Ln7 = dtoFrame.Ln7.ValueOrZero()
	dest.Ln8 = dtoFrame.Ln8.ValueOrZero()
}
