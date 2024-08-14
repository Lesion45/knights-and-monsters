package game

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten"
	"log"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

var frames map[string]Frames
var frame int
var lastKey ebiten.Key
var prevKey ebiten.Key
var levelImage *ebiten.Image

type Units map[uuid.UUID]*Unit

type Game struct {
	World World
}

type World struct {
	ID uuid.UUID
	Units
}

type Unit struct {
	ID         uuid.UUID
	X          float64
	Y          float64
	SpriteName string
}

func NewGame() (*Game, error) {
	var err error
	frames, err = LoadResources()
	fmt.Println(frames)
	levelImage, err = prepareLevelImage()
	if err != nil {
		log.Fatal(err)
	}

	return &Game{}, nil
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

// Update updates the current game state.
func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

// Draw draws the current game to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(levelImage, op)
}

func prepareLevelImage() (*ebiten.Image, error) {
	tileSize := 16
	level := LoadLevel()
	width := len(level[0])
	height := len(level)
	levelImage, err := ebiten.NewImage(width*tileSize, height*tileSize, ebiten.FilterDefault)
	if err != nil {
		return nil, err
	}

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(i*tileSize), float64(j*tileSize))

			img, err := ebiten.NewImageFromImage(frames[level[j][i]].Frames[0], ebiten.FilterDefault)
			if err != nil {
				return nil, err
			}
			levelImage.DrawImage(img, op)
		}
	}

	return levelImage, nil
}
