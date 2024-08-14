package main

import (
	"github.com/hajimehoshi/ebiten"
	kna "knights-and-monsters/game" // kna == Knights and Monsters
	"log"
)

func main() {
	game, err := kna.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(kna.ScreenWidth, kna.ScreenHeight)
	ebiten.SetWindowTitle("Knights and Monsters beta")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
