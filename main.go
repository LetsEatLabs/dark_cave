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
	screenWidth         = 640
	screenHeight        = 480
	textWidth           = 70 // Columns
	textAreaX           = 40
	textAreaY           = 40
	playerTextAreaLineY = 420
	playerTextAreaX     = 40
	playerTextAreaY     = 440
	wrapDistance        = 8 // Text wrap check distance
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
	isDebug         bool
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
	ebitenutil.DrawLine(screen, 0, playerTextAreaLineY, screenWidth, playerTextAreaLineY, color.White)
	text.Draw(screen, g.playerInputText, mplusNormalFont, playerTextAreaX, playerTextAreaY, color.White)

	// Text Display
	text.Draw(screen, g.text, mplusNormalFont, textAreaX, textAreaY, color.White)
}

func main() {
	g := &Game{}
	g.isDebug = false // Debug mode is off by default. To turn on, type "debug on"

	// Set up text area
	g.ps1 = "> "
	g.playerInputText = g.ps1
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Dark Cave")

	// Load the narrative
	loadGameLocations(g)
	loadObjectDetails(g)
	g.currentLocation = "cave_entrance"

	// Starting text
	g.text = "You realize that you have been having a day dream for a few seconds. \n\nType 'look' to look around. Type 'help' for more."

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

}
