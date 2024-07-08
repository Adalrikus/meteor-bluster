package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := NewGame()

	log.Println("Running the game.")
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
