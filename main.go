package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &Game{
		player:           NewPlayer(),
		meteorSpawnTimer: NewTimer(time.Millisecond * 5),
		meteors:          []*Meteor{},
	}

	log.Println("Running the game.")
	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
