package main

// Imports
import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Global constants
const (
	screenWidth  = 640
	screenHeight = 480
)

// Global Vars
var (
	mplusNormalFont font.Face
)

// Game Struct type
type Game struct {
	runes           []rune
	text            string
	playerInputText string
	counter         int
}

// Initialize some things

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		log.Fatal("Font is broke")
	}

	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    16,
		DPI:     72,
		Hinting: font.HintingFull,
	})
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

	text.Draw(screen, g.text, mplusNormalFont, 40, 40, color.White)
}

func main() {
	g := &Game{}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Test!")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}