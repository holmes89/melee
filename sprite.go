package main

import (
	"bytes"
	_ "embed"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

const SpriteType ComponentType = "SPRITE"

// SpriteComponent holds a reference to an image to be drawn
type SpriteComponent struct {
	image *ebiten.Image
}

func (t *SpriteComponent) Type() ComponentType { return SpriteType }

type SpriteRenderSystem struct {
	registry *Registry
}

func (s *SpriteRenderSystem) Draw(screen *ebiten.Image) {

	for _, e := range s.registry.Query(TransformType, SpriteType) {

		position := e.GetComponent(TransformType).(*TransformComponent)
		sprite := e.GetComponent(SpriteType).(*SpriteComponent)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(position.PosX, position.PosY)
		screen.DrawImage(sprite.image, op)
	}
}

// Dwarf

//go:embed assets/dwarf.png
var dwarfImage []byte

func NewDwarf() ComponentTyper {
	img, _, err := image.Decode(bytes.NewReader(dwarfImage))
	if err != nil {
		panic(err)
	}
	component := SpriteComponent{
		image: ebiten.NewImageFromImage(img),
	}
	return &component
}
