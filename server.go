package main

import (
	"log"
	"main/core"
	"main/services/cryptome"
)

func main() {
	app := core.Server()
	log.Printf("main.app: %p", app)
	cryptome.Route()
	app.Logger.Fatal(app.Start(":8080"))
}
