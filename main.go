package main

import (
	"github.com/xuedi/starraid/server"
)

func main() {
	app := server.Server{}
	app.Init()

	for {
		app.Tick()
	}
}
