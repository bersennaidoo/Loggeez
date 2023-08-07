package app

import (
	"github.com/bersennaidoo/Loggeez/internal/service/repo"
)

type App struct {
	Log *repo.Log
}

func NewApp(l *repo.Log) *App {
	return &App{
		Log: l,
	}
}
