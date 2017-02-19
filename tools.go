package main

import termbox "github.com/nsf/termbox-go"

func fill(x, y, w int, cell termbox.Cell) {
	for i := 0; i < w; i++ {
		termbox.SetCell(x+i, y, cell.Ch, cell.Fg, cell.Bg)
	}
}
