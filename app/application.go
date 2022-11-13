package app

import (
	"log"

	"github.com/audryus/ubiquitous-spoon/controller"
)

func Init() {
	handlers := controller.NewHandler()
	app := handlers.Init()
	log.Fatal(app.Listen(":8080"))
}
