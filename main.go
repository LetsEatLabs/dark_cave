package main

// Imports
import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

// Global constants
const (
	screenWidth  = 640
	screenHeight = 480
)

// Game Struct type
type Game struct{}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// pass
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Test!")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
