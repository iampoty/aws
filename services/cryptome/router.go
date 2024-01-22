package cryptome

import (
	"log"
	"main/core"
)

type (
	Control struct{}
)

var (
	controller Control
)

func Route() {
	app := core.Server()
	log.Printf("router.app: %p", app)
	app.GET("/", controller.hello)
}
