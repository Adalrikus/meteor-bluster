package main

import (
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	game          *Game
	shootCooldown *Timer
	sprite        *ebiten.Image
	position      Vector
	rotation      float64
}

func NewPlayer(g *Game) *Player {
	sprite := PlayerSprite

	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	return &Player{
		game:          g,
		shootCooldown: NewTimer(time.Millisecond),
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

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		p.shootCooldown.Reset()
		bulletSpawnOffset := 50.0

		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := Vector{
			X: p.position.X + halfW + math.Sin(p.rotation)*bulletSpawnOffset,
			Y: p.position.Y + halfH + math.Cos(p.rotation)*-bulletSpawnOffset,
		}

		bullet := NewBullet(spawnPos, p.rotation)
		p.game.AddBullet(bullet)
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

func (p *Player) Collider() Rect {
	bounds := p.sprite.Bounds()

	return Rect{
		X:      p.position.X,
		Y:      p.position.Y,
		Width:  float64(bounds.Dx()),
		Height: float64(bounds.Dy()),
	}
}
