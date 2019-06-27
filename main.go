package main

import (
	"fmt"

	"slack-sl/app"
)

func main() {
	s := app.NewServer()

	fmt.Printf("Hello World: %v", s)
}
