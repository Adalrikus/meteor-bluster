package main

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets/*
var assets embed.FS

var (
	PlayerSprite  = mustLoadImage("assets/PNG/Sprites/Ships/spaceShips_001.png")
	MeteorSprites = mustLoadImages("assets/PNG/Sprites/Meteors")
	LaserSprite   = mustLoadImage("assets/PNG/Sprites/Missiles/spaceMissiles_027.png")
)

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	log.Printf("Loading sprite: %s\n", name)
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Asset loaded.")

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(name string) []*ebiten.Image {
	files, err := assets.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	images := make([]*ebiten.Image, len(files))
	for i, file := range files {
		images[i] = mustLoadImage(name + "/" + file.Name())
	}

	return images
}
