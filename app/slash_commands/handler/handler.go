package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ojiry/slack-sl/app/slash_commands"
	"github.com/ojiry/slack-sl/app/slash_commands/store"
)

type SlashCommandsResponse struct {
	Type       string `json:"type"`
	Text       Text   `json:"text"`
	Value      string `json:"value"`
	CallbackID string `json:"callback_id"`
}

type Text struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

func SlashCommandsHandler(w http.ResponseWriter, req *http.Request) {
	uc := slash_commands.NewSlashCommandsUsecase(store.NewLunchStore())

	lunch, err := uc.CreateLunch(parseSlashCommandsParameter())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := SlashCommandsResponse{
		Type: "button",
		Text: Text{
			Type:  "plain_text",
			Value: "Click Me",
		},
		Value:      "click_me_123",
		CallbackID: lunch.ID,
	}

	json, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}

func parseSlashCommandsParameter() *slash_commands.SlashCommands {
	return &slash_commands.SlashCommands{}
}
