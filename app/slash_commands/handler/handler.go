package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type SlashCommandsResponse struct {
	Type  string
	Text  Text
	Value string
}

type Text struct {
	Type  string
	Value string
}

func SlashCommandsHandler(w http.ResponseWriter, req *http.Request) {
	res := SlashCommandsResponse{
		Type: "button",
		Text: Text{
			Type:  "plain_text",
			Value: "Click Me",
		},
		Value: "click_me_123",
	}

	json, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(json))
}
