package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth      = 500
	screenHeight     = 500
	characterBoundry = 350
)

func main() {

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Melee")
	if err := ebiten.RunGame(&Game{
		Player: NewDwarf(),
		Level:  NewForestLevel(),
	}); err != nil {
		log.Fatal(err)
	}
}
