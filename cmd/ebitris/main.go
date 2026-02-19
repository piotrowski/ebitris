package main

import "github.com/piotrowski/ebitris/internal/game"

func main() {
	err := game.Start()
	if err != nil {
		panic(err)
	}
}
