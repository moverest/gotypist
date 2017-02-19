package main

import (
	"strings"

	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

// Line represent a line in the text.
type Line struct {
	Txt string
}

// Draw draws the content of the line
func (l *Line) Draw(x, y, w, cursorPos int, correct bool) {
	fill(x, y, w,
		termbox.Cell{Ch: ' ', Bg: defaultLineBgColor, Fg: defaultLineFgColor})

	pos := 0
	for i, c := range l.Txt {
		fgColor := defaultLineFgColor
		bgColor := defaultLineBgColor

		if i == cursorPos {
			if correct {
				fgColor = defaultLineFgCursorColor
				bgColor = defaultLineBgCursorColor
			} else {
				fgColor = defaultLineFgCursorErrColor
				bgColor = defaultLineBgCursorErrColor
			}
		}

		termbox.SetCell(x+pos, y, c, fgColor, bgColor)
		pos += runewidth.RuneWidth(c)
	}
}

// Correct returns true if the beginning correspond to the given string.
func (l *Line) Correct(s string) bool {
	if len(s) > len(l.Txt) {
		return false
	}

	return l.Txt[:len(s)] == s
}

// Done returns true if the given string contains the whole line.
func (l *Line) Done(s string) bool {
	return len(s) >= len(l.Txt)+1 && l.Txt+" " == s[:len(l.Txt)+1]
}

// Text represent a text
type Text struct {
	Lines       []Line
	CurrentLine int
}

// Draw draw the whole text.
func (txt *Text) Draw(x, y, w, cursorPos int, correct bool) {
	for i := range txt.Lines {
		lineCursorPos := -1
		if i == txt.CurrentLine {
			lineCursorPos = cursorPos
		}

		txt.Lines[i].Draw(x, y+i, w, lineCursorPos, correct)
	}
}

// LoadText load a string to be used as the text.
func (txt *Text) LoadText(s string) {
	lines := strings.Split(s, "\n")

	txt.CurrentLine = 0
	txt.Lines = make([]Line, len(lines))

	for i := range lines {
		txt.Lines[i].Txt = lines[i]
	}
}

// Correct returns true if the beginning of the current line correspond to the
// given string.
func (txt *Text) Correct(s string) bool {
	return txt.Lines[txt.CurrentLine].Correct(s)
}

// Done advance the game by one line if given string correspond to the current
// one. Returns true if the game is ended.
func (txt *Text) Done(s string) (done bool, advancedLine bool) {
	if txt.Lines[txt.CurrentLine].Done(s) {
		txt.CurrentLine++
		done = txt.CurrentLine >= len(txt.Lines)
		advancedLine = true
		return
	}

	return
}

// NewTextFromString creates a text from a string.
func NewTextFromString(s string) *Text {
	txt := &Text{}
	txt.LoadText(s)

	return txt
}
