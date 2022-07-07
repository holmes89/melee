package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	spriteRenderSystem *SpriteRenderSystem
}

func (a *Game) Update() error { return nil }
func (a *Game) Draw(screen *ebiten.Image) {
	a.spriteRenderSystem.Draw(screen)
}
func (a *Game) Layout(ow, oh int) (int, int) { return ow, oh }
