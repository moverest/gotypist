package main

import termbox "github.com/nsf/termbox-go"

// Game represents the game.
type Game struct {
	Input *InputBox
	Txt   *Text
}

// NewGame creates a new game from given text.
func NewGame(txt string) *Game {
	return &Game{
		Input: &InputBox{},
		Txt:   NewTextFromString(txt),
	}
}

// Draw draws the game.
func (g *Game) Draw(x, y, w int) {
	g.Txt.Draw(x, y, w, g.Input.CursorPosition(), g.Txt.Correct(g.Input.Txt))
	g.Input.Draw(x, y+1+len(g.Txt.Lines), w)

	termbox.Flush()
}

// LetterPress process a letter key pressed event.
func (g *Game) LetterPress(key rune) {
	g.Input.Append(key)
}

// EnterKeyPress process an enter key being pressed
func (g *Game) EnterKeyPress() {
	g.Input.Txt = ""
}

// Done advance the game if current line is done. If the game is done, it
// returns true.
func (g *Game) Done() bool {
	done, advancedLine := g.Txt.Done(g.Input.Txt)
	if advancedLine {
		g.Input.Txt = ""
	}

	return done
}
