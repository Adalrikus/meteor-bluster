package main

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	sprite   *ebiten.Image
	position Vector
	rotation float64
}

func NewPlayer() *Player {
	sprite := PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	return &Player{
		position: Vector{
			X: ScreenWidth/2 - halfW,
			Y: ScreenHeight/2 - halfH,
		},
		sprite: sprite,
	}
}

func (p *Player) Update() error {
	speed := float64(300 / ebiten.TPS())
	rotationalSpeed := math.Pi / float64(ebiten.TPS())

	var delta Vector
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		delta.X = math.Sin(p.rotation) * speed
		delta.Y = math.Cos(p.rotation) * -speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyJ) {
		delta.X = math.Sin(p.rotation) * -speed
		delta.Y = math.Cos(p.rotation) * speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		p.rotation += rotationalSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) {
		p.rotation -= rotationalSpeed
	}

	p.position.X += delta.X
	p.position.Y += delta.Y

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
