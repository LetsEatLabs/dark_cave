package main

// Imports
import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Global constants
const (
	screenWidth  = 640
	screenHeight = 480
)

// Game Struct type
type Game struct {
	runes   []rune
	text    string
	counter int
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {

	g.runes = ebiten.AppendInputChars(g.runes[:0])
	g.text += string(g.runes)
	g.counter++

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	t := g.text

	// Blink cursor
	if g.counter%60 < 30 {
		t += "-"
	}

	ebitenutil.DebugPrint(screen, t)
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Test!")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
