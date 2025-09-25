package backend

import (
	"context"
	"fmt"
)

type App struct {
	ctx  context.Context
	Game *Game
}

func NewApp() *App {
	return &App{
		Game: &Game{},
	}
}

func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
