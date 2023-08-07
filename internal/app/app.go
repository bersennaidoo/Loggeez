package app

import (
	r "github.com/bersennaidoo/Loggeez/internal/service/repo"
)

type App struct {
	Log *r.Log
}

func NewApp(l *r.Log) *App {
	return &App{
		Log: l,
	}
}
