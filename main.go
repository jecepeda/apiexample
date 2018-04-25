package main

import (
	"log"

	"github.com/jcepedavillamayor/apiexample/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
