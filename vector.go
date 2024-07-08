package main

import "math"

type Vector struct {
	X float64
	Y float64
}

func (v *Vector) Normalize() Vector {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)

	return Vector{
		X: v.X / length,
		Y: v.Y / length,
	}
}
