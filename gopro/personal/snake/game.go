package main

import (
	"time"
)

type Game struct {
	Screen    tcell.Screen
	snakeBody SnakeBody
}

func (g *Game) Run() {

	defStyle := tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite)
	g.Screen.SetStyle(defStyle)
	width, height := g.Screen.Size()
	snakeStyle := tcell.StyleDefault.Background(tcell.ColorWhite).Foreground(tcell.ColorWhite)

	for {
		g.Screen.Clear()
		g.snakeBody.Update(width, height)
		g.Screen.SetContent(g.snakeBody.X, g.snakeBody.Y, ' ', nil, snakeStyle)
		time.Sleep(40 * time.Millisecond)
		g.Screen.Show()
	}

}
