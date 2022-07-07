package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 500
	screenHeight = 500
)

func main() {

	// entities
	registry := Registry{}
	gopher := registry.NewEntity()

	// components
	gopher.AddComponent(&TransformComponent{PosY: 200, PosX: 200})
	gopher.AddComponent(NewDwarf())

	// systems
	spriteRenderSystem := &SpriteRenderSystem{registry: &registry}

	// the game implementing the ebiten.Game interface

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Melee")
	if err := ebiten.RunGame(&Game{
		spriteRenderSystem: spriteRenderSystem,
	}); err != nil {
		log.Fatal(err)
	}
}
