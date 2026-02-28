package game

import (
	"fmt"
	"log/slog"
	"os"

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

func Start() error {
	level := slog.LevelInfo
	if os.Getenv("LOG_LEVEL") == "debug" {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: level,
	})))

	ebiten.SetWindowSize(600, 800)
	ebiten.SetWindowTitle("Ebitris")

	manager := scene.NewManager()

	if err := ebiten.RunGame(&Game{
		manager: manager,
	}); err != nil {
		return fmt.Errorf("failed to run game: %w", err)
	}
	return nil
}
