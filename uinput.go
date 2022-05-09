package main

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func MoveInputToTerminal(g *Game) {
	g.text += fmt.Sprintf("You: %s\n", strings.Replace(g.playerInputText, "> ", "", 1))
	g.playerInputText = g.ps1
}

func RepeatKeyPressed(key ebiten.Key) {
	// pass
}
