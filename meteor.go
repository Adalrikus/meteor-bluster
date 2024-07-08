package main

import (
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	sprite        *ebiten.Image
	position      Vector
	movement      Vector
	rotation      float64
	rotationSpeed float64
}

func NewMeteor() *Meteor {
	sprite := MeteorSprites[rand.Intn(len(MeteorSprites))]

	// Get a center of the screen as a target
	target := Vector{
		X: ScreenWidth / 2,
		Y: ScreenHeight / 2,
	}

	// Get a random angle from 0 to 2 * Pi
	theta := rand.Float64() * 2 * math.Pi

	// ScreenWidth is our diameter, hence radius is half the size
	radius := ScreenWidth / 2.0

	pos := Vector{
		X: radius*math.Cos(theta) + target.X,
		Y: radius*math.Sin(theta) + target.Y,
	}

	// Randomized velocity of the meteor
	velocity := 0.25 + rand.Float64()*1.5

	// Direction of the meteor
	direction := Vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	// Normalize the vector — get just the direction without the length
	normalDirect := direction.Normalize()

	movement := Vector{
		X: normalDirect.X * velocity,
		Y: normalDirect.Y * velocity,
	}

	log.Println("Meteor created.")

	return &Meteor{
		position:      pos,
		movement:      movement,
		sprite:        sprite,
		rotationSpeed: -0.02 + rand.Float64()*0.04,
	}
}

func (m *Meteor) Update() {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}
