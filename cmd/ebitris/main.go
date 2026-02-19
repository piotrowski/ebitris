package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/piotrowski/ebitris/internal/scene"
)

type Game struct {
	manager *scene.Manager
}

func (g *Game) Update() error {
	return g.manager.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.manager.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 600, 800
}

func main() {
	ebiten.SetWindowSize(600, 800)
	ebiten.SetWindowTitle("Ebitris")

	manager := scene.NewManager()
	manager.SwitchTo(scene.NewGameplayScene(manager, 10, 20))

	if err := ebiten.RunGame(&Game{
		manager: manager,
	}); err != nil {
		log.Fatal(err)
	}
}
