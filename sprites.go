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
)

func mustLoadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Assets loaded.")

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(name string) []*ebiten.Image {
	files, err := assets.ReadDir(name)
	if err != nil {
		log.Fatal(err)
	}

	var images []*ebiten.Image
	for _, file := range files {
		log.Printf("Loading sprite: %s\n", file.Name())
		content, err := assets.Open(name + "/" + file.Name())
		if err != nil {
			log.Fatal(err)
		}

		img, _, err := image.Decode(content)
		if err != nil {
			log.Fatal(err)
		}

		image := ebiten.NewImageFromImage(img)
		images = append(images, image)
	}

	log.Println("Assets loaded.")

	return images
}
