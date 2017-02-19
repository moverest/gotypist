package main

import (
	"io/ioutil"
	"os"

	"github.com/nsf/termbox-go"
)

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

	var game *Game

	if len(os.Args) > 1 {
		b, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			panic("Could not load given file.")
		}

		game = NewGame(string(b))
	} else {
		game = NewGame(`Gotypist is a simple typing training program written in go.

To quit hit the ESCAPE key.

To load a file as training text,
Pass it as an argument.

You cannot use backspace to correct a mistake.
Hit ENTER to retry the current line.`)
	}

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
