package main

import (
	"log"

	"github.com/markbates/gobular/actions"
)

func main() {
	app := actions.App()
	log.Fatal(app.Serve())
}
