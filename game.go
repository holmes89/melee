package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	Player Character
	input  Input
	Level  Level
}

func (a *Game) Update() error {
	a.input.Update()

	a.Player.SetAttacking(a.input.IsAttackJustPressed() > 0)
	a.Player.SetJumping(a.input.StateForUp() > 0)

	moving := a.input.StateForLeft() > 0 || a.input.StateForRight() > 0
	a.Player.SetMoving(moving)
	if moving {
		a.Player.SetFacingRight(a.input.StateForRight() > 0)
	}
	a.Level.SetMoving(moving && a.Player.AtBoundary())

	a.Level.Update()
	a.Player.Update()
	return nil
}
func (a *Game) Draw(screen *ebiten.Image) {
	a.Level.Draw(screen)
	a.Player.Draw(screen)
}
func (a *Game) Layout(ow, oh int) (int, int) { return ow, oh }

type GameComponent interface {
	Update() error
	Draw(screen *ebiten.Image)
}
