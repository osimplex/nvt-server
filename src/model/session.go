package model

type Session struct {
	ClientAddr        string
	Transaction       string
	Input             string
	SearchType        string
	Beep              string
	FrameName         string
	PreviousFrameName string

	Register   SessionRegister
	FrameLines SessionFrameLine
}
