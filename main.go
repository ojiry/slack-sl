package main

import (
	"log"

	"github.com/ojiry/slack-sl/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
