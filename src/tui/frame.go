package tui

import (
	"errors"
	"fmt"
	"nvt-server/src/model"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	ClearScreenSequence    = "\x1b\x5b\x32\x4a"
	cursorPositioningRune  = "#"
	cursorPositioningFrame = "\x1b\x5b%d\x3b%d\x48"
)

var (
	ErrNoPositioningRune = errors.New("no positioning rune found in string")
)

func DetermineCursorPosition(frame model.Frame) (string, error) {
	positionFound := false

	lines := []string{
		frame.Ln1,
		frame.Ln2,
		frame.Ln3,
		frame.Ln4,
		frame.Ln5,
		frame.Ln6,
		frame.Ln7,
		frame.Ln8,
	}

	position := [2]int{8, utf8.RuneCountInString(lines[7])}

	for i, label := range lines {
		if j := strings.Index(label, cursorPositioningRune); j != -1 {
			position[0] = i + 1
			position[1] = j + 1
			positionFound = true
			break
		}
	}

	if !positionFound {
		return cursorPositionVT100(position), ErrNoPositioningRune
	}

	return cursorPositionVT100(position), nil
}

func cursorPositionVT100(position [2]int) string {
	return fmt.Sprintf(cursorPositioningFrame, position[0], position[1])
}

func BuildFrameString(frame model.Frame, frameLine model.SessionFrameLine) string {
	sbFrame := strings.Builder{}
	sbFrame.WriteString(frame.Ln1 + frameLine.Rl1)
	sbFrame.WriteString("\n\r")
	sbFrame.WriteString(frame.Ln2 + frameLine.Rl2)
	sbFrame.WriteString("\n\r")
	sbFrame.WriteString(frame.Ln3 + frameLine.Rl3)
	sbFrame.WriteString("\n\r")
	sbFrame.WriteString(frame.Ln4 + frameLine.Rl4)
	sbFrame.WriteString("\n\r")
	sbFrame.WriteString(frame.Ln5 + frameLine.Rl5)
	sbFrame.WriteString("\n\r")
	sbFrame.WriteString(frame.Ln6 + frameLine.Rl6)
	sbFrame.WriteString("\n\r")
	sbFrame.WriteString(frame.Ln7 + frameLine.Rl7)
	sbFrame.WriteString("\n\r")
	sbFrame.WriteString(frame.Ln8 + frameLine.Rl8)

	expression := regexp.MustCompile(cursorPositioningRune + "N?")
	result := expression.ReplaceAllString(sbFrame.String(), "")

	return result
}

func MessageOfTheDay() string {
	date := fmt.Sprint(time.Now().Format("02/01/2006"))

	sbMotd := strings.Builder{}
	sbMotd.WriteString("-------------------\n\r")
	sbMotd.WriteString("                   \n\r")
	sbMotd.WriteString(" -Network Virtual- \n\r")
	sbMotd.WriteString(" -Terminal System- \n\r")
	sbMotd.WriteString("                   \n\r")
	sbMotd.WriteString("  Date: ")
	sbMotd.WriteString(date + "\n\r")
	sbMotd.WriteString("                    \n\r")
	sbMotd.WriteString("-------------------")

	return sbMotd.String()
}
