package main

import (
	"bytes"
	_ "embed"
	"image"
	"math"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadSpritesheet(input []byte, row int, n int, width int, height int) []*ebiten.Image {
	var sprites []*ebiten.Image

	spritesheet, _, _ := image.Decode(bytes.NewReader(input))
	ebitenImage := ebiten.NewImageFromImage(spritesheet)

	y0 := row * height
	for i := 0; i < n; i++ {
		dimensions := image.Rect(i*width, y0, (i+1)*width, y0+height)
		sprite := ebitenImage.SubImage(dimensions).(*ebiten.Image)
		sprites = append(sprites, sprite)
	}

	return sprites
}

type characterState string

const (
	Idle       characterState = "idle"
	Walk       characterState = "walk"
	JumpAttack characterState = "jumpattack"
	Attack     characterState = "attack"
	Spin       characterState = "spin"
	Jump       characterState = "jump"
	Damage     characterState = "damage"
	Death      characterState = "death"
)

type characterSpriteState struct {
	Left  map[characterState][]*ebiten.Image
	Right map[characterState][]*ebiten.Image
}

type Character interface {
	GameComponent
	SetAttacking(bool)
	SetJumping(bool)
	SetFacingRight(bool)
	SetMoving(bool)
	AtBoundary() bool
}

// Dwarf

//go:embed assets/dwarf.png
var dwarfImage []byte

type dwarf struct {
	image             *ebiten.Image
	frames            []*ebiten.Image
	spriteStates      characterSpriteState
	state             characterState
	posX              float64
	posY              float64
	currentFrameIndex int
	count             float64
	animationSpeed    float64
	facingRight       bool
	jumping           bool
	attacking         bool
	moving            bool
}

func NewDwarf() Character {
	state := characterSpriteState{}

	state.Left = make(map[characterState][]*ebiten.Image)
	state.Left[Idle] = LoadSpritesheet(dwarfImage, 8, 5, 38, 32)
	state.Left[Walk] = LoadSpritesheet(dwarfImage, 9, 8, 38, 32)
	state.Left[JumpAttack] = LoadSpritesheet(dwarfImage, 10, 7, 38, 32)
	state.Left[Attack] = LoadSpritesheet(dwarfImage, 11, 6, 38, 32)
	state.Left[Spin] = LoadSpritesheet(dwarfImage, 12, 2, 38, 32)
	state.Left[Jump] = LoadSpritesheet(dwarfImage, 13, 5, 38, 32)
	state.Left[Damage] = LoadSpritesheet(dwarfImage, 14, 4, 38, 32)
	state.Left[Death] = LoadSpritesheet(dwarfImage, 15, 7, 38, 32)

	state.Right = make(map[characterState][]*ebiten.Image)
	state.Right[Idle] = LoadSpritesheet(dwarfImage, 0, 5, 38, 32)
	state.Right[Walk] = LoadSpritesheet(dwarfImage, 1, 8, 38, 32)
	state.Right[JumpAttack] = LoadSpritesheet(dwarfImage, 2, 7, 38, 32)
	state.Right[Attack] = LoadSpritesheet(dwarfImage, 3, 6, 38, 32)
	state.Right[Spin] = LoadSpritesheet(dwarfImage, 4, 2, 38, 32)
	state.Right[Jump] = LoadSpritesheet(dwarfImage, 5, 5, 38, 32)
	state.Right[Damage] = LoadSpritesheet(dwarfImage, 6, 4, 38, 32)
	state.Right[Death] = LoadSpritesheet(dwarfImage, 7, 7, 38, 32)

	frames := state.Right[Idle]
	return &dwarf{
		frames:         frames,
		image:          frames[0],
		animationSpeed: .125,
		spriteStates:   state,
		state:          Idle,
		posX:           40,
		posY:           450,
		facingRight:    true,
	}
}

func (d *dwarf) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(d.posX, d.posY)
	screen.DrawImage(d.image, op)
}

func (d *dwarf) AtBoundary() bool {
	return d.posX >= characterBoundry
}

func (d *dwarf) Update() error {

	// Find current state
	state := Idle
	if d.moving {
		state = Walk
	}
	if d.jumping {
		state = Jump
	}
	if d.attacking {
		state = Attack
	}
	if d.jumping && d.attacking {
		state = JumpAttack
	}
	// advance animation
	d.count += d.animationSpeed
	d.currentFrameIndex = int(math.Floor(d.count))

	// Update State
	if state != d.state {
		d.state = state
		d.count = 0
		d.currentFrameIndex = 0
	}

	// Update direction
	if d.facingRight {
		d.frames = d.spriteStates.Right[d.state]
		if d.moving && d.posX <= characterBoundry {
			d.posX += 1
		}
	} else {
		d.frames = d.spriteStates.Left[d.state]
		if d.moving && d.posX > 0 {
			d.posX -= 1
		}
	}

	if d.currentFrameIndex >= len(d.frames) { // restart animation
		d.count = 0
		d.currentFrameIndex = 0
	}

	// update image
	d.image = d.frames[d.currentFrameIndex]
	return nil
}
func (d *dwarf) SetAttacking(attacking bool) {
	d.attacking = attacking
}
func (d *dwarf) SetJumping(jumping bool) {
	d.jumping = jumping
}
func (d *dwarf) SetFacingRight(facingRight bool) {
	d.facingRight = facingRight
}
func (d *dwarf) SetMoving(moving bool) {
	d.moving = moving
}
