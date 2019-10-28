package app

import (
	"net/http"

	scHandler "github.com/ojiry/slack-sl/app/slash_commands/handler"
)

func Run() error {
	http.HandleFunc("/slash_commands", scHandler.SlashCommandsHandler)
	return http.ListenAndServe(":8080", nil)
}
