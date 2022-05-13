package main

// Imports
import (
	"image/color"
	"log"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"

	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Global constants
const (
	screenWidth  = 640
	screenHeight = 480
	textWidth    = 70 // Columns
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
	ps1             string
	counter         int
	locations       []Location
	currentLocation string
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

	if err != nil {
		log.Fatal(err)
	}
}

/////

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) Update() error {

	g.runes = ebiten.AppendInputChars(g.runes[:0])
	g.playerInputText += string(g.runes)
	g.counter++

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		command := ParseUserCommand(g.playerInputText)
		MoveInputToTerminal(g)
		HandleCommand(g, command)

	}

	if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
		if len(g.playerInputText) > 2 {
			g.playerInputText = g.playerInputText[:len(g.playerInputText)-1]
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	t := g.text

	// Blink cursor
	if g.counter%60 < 30 {
		t += "-"
	}

	// Player input area
	ebitenutil.DrawLine(screen, 0, 420, screenWidth, 420, color.White)
	text.Draw(screen, g.playerInputText, mplusNormalFont, 40, 440, color.White)

	// Text Display
	text.Draw(screen, g.text, mplusNormalFont, 40, 40, color.White)
}

func main() {
	g := &Game{}

	// Set up text area
	g.ps1 = "> "
	g.playerInputText = g.ps1
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dark Cave")

	// Load the narrative
	loadGameLocations(g)
	loadObjectDetails(g)
	g.currentLocation = "dark_cave_entrance"

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
