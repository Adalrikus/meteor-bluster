package main

import (
	"embed"
	"image"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed assets/*
var assets embed.FS

var (
	PlayerSprite  = mustLoadImage("assets/PNG/Sprites/Ships/spaceShips_001.png")
	MeteorSprites = mustLoadImages("assets/PNG/Sprites/Meteors")
	LaserSprite   = mustLoadImage("assets/PNG/Sprites/Missiles/spaceMissiles_027.png")
	ScoreFont     = mustLoadFont("assets/Fonts/KenneyMiniSquareMono.ttf")
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

func mustLoadFont(name string) *text.GoXFace {
	f, err := assets.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		log.Fatal(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}

	goXFace := text.NewGoXFace(face)

	return goXFace
}
