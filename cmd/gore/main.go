package main

import (
	"log"

	"github.com/m3tro1d/gore/pkg/gore/app"
	"github.com/m3tro1d/gore/pkg/gore/infrastructure"
)

func main() {
	renamer := app.NewRenamer(
		infrastructure.NewFileLister(),
		infrastructure.NewEditor(),
		app.NewSanityChecker(),
	)

	if err := renamer.Rename("."); err != nil {
		log.Fatal(err)
	}
}
