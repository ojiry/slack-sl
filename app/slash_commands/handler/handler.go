package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/ojiry/slack-sl/app/slash_commands"
	"github.com/ojiry/slack-sl/app/slash_commands/store"
)

func SlashCommandsHandler(w http.ResponseWriter, req *http.Request) {
	sc := makeSlashCommands(req)
	u := slash_commands.NewSlashCommandsUsecase(store.NewLunchStore())
	l, err := u.CreateLunch(*sc)
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Printf("%v\n", l)
	io.WriteString(w, "Hello, world!\n")
}

func makeSlashCommands(req *http.Request) *slash_commands.SlashCommands {
	return &slash_commands.SlashCommands{}
}
