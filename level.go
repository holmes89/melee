package main

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"

	_ "image/png"
)

type Level struct {
	background Background
	moving     bool
}

func (l *Level) Update() error {
	l.background.SetMoving(l.moving)
	return l.background.Update()
}
func (l *Level) Draw(screen *ebiten.Image) {
	l.background.Draw(screen)
}
func (l *Level) SetMoving(moving bool) {
	l.moving = moving
}

// Forest

//go:embed assets/backgrounds/forest/layers/0010.png
var forestBackground []byte

//go:embed assets/backgrounds/forest/layers/0009.png
var forestBackground2 []byte

//go:embed assets/backgrounds/forest/layers/0000.png
var forestGround0 []byte // TODO foreground

//go:embed assets/backgrounds/forest/layers/0001.png
var forestGround1 []byte

//go:embed assets/backgrounds/forest/layers/0002.png
var forestCannopy []byte

//go:embed assets/backgrounds/forest/layers/0003.png
var forestTrunks1 []byte

//go:embed assets/backgrounds/forest/layers/0005.png
var forestTrunks2 []byte

//go:embed assets/backgrounds/forest/layers/0006.png
var forestTrunks3 []byte

//go:embed assets/backgrounds/forest/layers/0008.png
var forestTrunks4 []byte

//go:embed assets/backgrounds/forest/layers/0004.png
var forestLight1 []byte

//go:embed assets/backgrounds/forest/layers/0007.png
var forestLight2 []byte

func NewForestLevel() Level {
	return Level{
		background: Background{
			layers: []*Layer{
				NewLayer(forestBackground, 0),
				NewLayer(forestBackground2, 0),
				NewLayer(forestTrunks4, 1),
				NewLayer(forestTrunks3, 2),
				NewLayer(forestLight2, 2),
				NewLayer(forestTrunks2, 3),
				NewLayer(forestLight1, 3),
				NewLayer(forestTrunks1, 4),
				NewLayer(forestGround1, 4),
				NewLayer(forestGround0, 4),
				NewLayer(forestCannopy, 4),
			},
		},
	}
}
