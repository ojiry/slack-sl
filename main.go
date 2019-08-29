package main

import (
	"fmt"

	"github.com/ojiry/slack-sl/app"
)

func main() {
	s := app.NewServer()

	fmt.Printf("Hello World: %v", s)
}
