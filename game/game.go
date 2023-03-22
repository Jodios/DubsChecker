package game

import (
	"image/color"
	"math/rand"
	"os"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/jodios/dubschecker/assets/fonts"
	"github.com/jodios/dubschecker/utils"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

const (
	screenWidth  = 800
	screenHeight = 700
)

const (
	// Gameboy resolution
	innerWidth  = 160
	innerHeight = 140
)

type Game struct {
	checkemButton    *Button
	Loader           *utils.Loader
	Font             font.Face
	Digits           int
	Background       *ebiten.Image
	BackgroundColour color.Color
	FontColour       color.Color
	IsDubs           bool
}

func NewGame(title string) *Game {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(title)
	loader := utils.NewLoader(os.DirFS("./assets"))
	checkem, err := loader.Image("checkem.png")
	if err != nil {
		panic(err)
	}
	btnIdle, err := loader.Image("button_idle.png")
	if err != nil {
		panic(err)
	}
	btnPressed, err := loader.Image("button.png")
	if err != nil {
		panic(err)
	}
	mf, err := opentype.Parse(fonts.Minecraft_ttf)
	if err != nil {
		panic(err)
	}
	font, err := opentype.NewFace(mf, &opentype.FaceOptions{
		Size:    20,
		DPI:     80,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}
	// 72649787
	// 10000000
	// 99999999

	return &Game{
		Loader:           loader,
		Font:             font,
		Digits:           rand.Intn(99999999-10000000) + 10000000,
		Background:       ebiten.NewImageFromImage(checkem),
		BackgroundColour: color.RGBA{213, 218, 241, 255},
		FontColour:       color.RGBA{38, 22, 139, 255},
		checkemButton: &Button{
			IdleButton: &Button{
				Image:         ebiten.NewImageFromImage(btnIdle),
				XPaddingLeft:  3,
				XPaddingRight: 3,
				X:             (innerWidth / 2) - (btnIdle.Bounds().Dx() / 2),
				Y:             (innerHeight / 2) - (btnIdle.Bounds().Dy() / 2) + 50,
			},
			PressedButton: &Button{
				Image:         ebiten.NewImageFromImage(btnPressed),
				XPaddingLeft:  3,
				XPaddingRight: 3,
				X:             (innerWidth / 2) - (btnPressed.Bounds().Dx() / 2),
				Y:             (innerHeight / 2) - (btnPressed.Bounds().Dy() / 2) + 50,
			},
		},
	}
}

func (g *Game) Start() {
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

func (g *Game) Update() error {
	// Update logical game state
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) && g.checkemButton.IdleButton.IsClicked(ebiten.CursorPosition()) {
		g.checkemButton.Clicked = true
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButton0) && g.checkemButton.Clicked {
		g.checkemButton.Clicks++
		g.checkemButton.Clicked = false
		g.Digits = rand.Intn(99999999-10000000) + 10000000
		g.IsDubs = checkdubs(g.Digits)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// update games rendering
	screen.Fill(color.RGBA{213, 218, 241, 255})
	if g.IsDubs {
		screen.DrawImage(g.Background, &ebiten.DrawImageOptions{})
	}
	g.checkemButton.ShowIdle(screen)
	text.Draw(screen, ">"+strconv.Itoa(g.Digits), g.Font, innerWidth/6, innerHeight/6, g.FontColour)
	if g.checkemButton.Clicked {
		g.checkemButton.Click(screen)
		text.Draw(screen, ">"+strconv.Itoa(g.Digits), g.Font, innerWidth/6, innerHeight/6, g.FontColour)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// screen is auto scaled
	return innerWidth, innerHeight
}

func checkdubs(digits int) bool {
	a := digits % 10
	digits /= 10
	b := digits % 10
	return a == b
}
