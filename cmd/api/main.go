package main

import (
	"github.com/bersennaidoo/Loggeez/internal/app"
	"github.com/bersennaidoo/Loggeez/internal/service/repo"
)

func main() {

	lg := repo.NewLog()
	app := app.NewApp(lg)
	app.RunAPI(":8080")
}
