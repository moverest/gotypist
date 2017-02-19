package main

import (
	runewidth "github.com/mattn/go-runewidth"
	termbox "github.com/nsf/termbox-go"
)

// InputBox represent a input box.
type InputBox struct {
	Txt string
}

// Draw draws a input box.
func (ib *InputBox) Draw(x, y, w int) {
	fill(x, y, w, termbox.Cell{
		Ch: ' ',
		Bg: defaultLineBgColor,
		Fg: defaultLineFgColor})

	pos := 0
	for _, c := range ib.Txt {
		if pos > w {
			termbox.SetCell(x+w, y, '+', defaultInputBoxFgColor,
				defaultInputBoxBgColor)
			continue
		}

		termbox.SetCell(x+pos, y, c, defaultInputBoxFgColor, defaultInputBoxBgColor)
		pos += runewidth.RuneWidth(c)
	}

	fill(x, y+1, w, termbox.Cell{
		Ch: '-',
		Bg: defaultInputBoxBgColor,
		Fg: defaultInputBoxFgColor,
	})

	termbox.SetCursor(x+pos, y)
}

// CursorPosition returns the cursor position.
func (ib *InputBox) CursorPosition() int {
	return len(ib.Txt)
}

// Append append a letter to the input box.
func (ib *InputBox) Append(letter rune) {
	ib.Txt += string(letter)
}
