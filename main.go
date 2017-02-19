package main

import "github.com/nsf/termbox-go"

const (
	defaultLineBgColor          = termbox.ColorDefault
	defaultLineFgColor          = termbox.ColorDefault
	defaultLineBgCursorColor    = termbox.ColorGreen
	defaultLineFgCursorColor    = termbox.ColorBlack
	defaultLineBgCursorErrColor = termbox.ColorRed
	defaultLineFgCursorErrColor = termbox.ColorWhite

	defaultInputBoxBgColor = termbox.ColorDefault
	defaultInputBoxFgColor = termbox.ColorDefault
)

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)

	width, _ := termbox.Size()

	game := NewGame("This is a test.\nA real test.")

	stop := false
	for !stop {
		game.Draw(0, 0, width)

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventInterrupt:
			stop = true
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				stop = true
			case termbox.KeySpace:
				game.LetterPress(' ')
			case termbox.KeyEnter:
				game.EnterKeyPress()

			default:
				if ev.Ch != 0 {
					game.LetterPress(ev.Ch)
				}
			}
		}

		if game.Done() {
			stop = true
		}
	}
}
