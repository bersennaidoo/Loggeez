package main

import (
	"github.com/bersennaidoo/Loggeez/internal/app"
	r "github.com/bersennaidoo/Loggeez/internal/service/repo"
)

func main() {

	lg := r.NewLog()
	app := app.NewApp(lg)
	app.RunAPI(":8080")
}
