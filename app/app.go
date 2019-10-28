package app

import (
	"log"
	"net/http"
	"os"

	scHandler "github.com/ojiry/slack-sl/app/slash_commands/handler"
)

func Run() error {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/slash_commands", scHandler.SlashCommandsHandler)

	return http.ListenAndServe(":"+port, nil)
}
