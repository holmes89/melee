package main

import (
	"bytes"
	_ "embed"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type Layer struct {
	image         *ebiten.Image
	posX          float64
	paralaxOffset float64
}

func NewLayer(input []byte, paralax int, dimensions ...image.Rectangle) *Layer {
	img, _, _ := image.Decode(bytes.NewReader(input))
	return &Layer{
		image:         ebiten.NewImageFromImage(img).SubImage(image.Rect(0, 250, 928, 1000)).(*ebiten.Image),
		paralaxOffset: float64(paralax) * .25,
	}
}

func (l *Layer) Update() error {
	l.posX -= l.paralaxOffset
	if (l.posX-50)+478 <= 0 {
		l.posX = 0
	}
	return nil
}

func (l *Layer) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(l.posX), 0)
	screen.DrawImage(l.image, op)
}

type Background struct {
	layers []*Layer
	moving bool
}

func (b *Background) Update() error {
	if !b.moving {
		return nil
	}
	for _, l := range b.layers {
		l.Update()
	}
	return nil
}

func (b *Background) Draw(screen *ebiten.Image) {
	for _, l := range b.layers {
		l.Draw(screen)
	}
}

func (b *Background) SetMoving(moving bool) {
	b.moving = moving
}
