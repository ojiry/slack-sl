package store

import (
	"fmt"

	"github.com/ojiry/slack-sl/app/slash_commands"
)

type LunchStore struct{}

func NewLunchStore() *LunchStore {
	return &LunchStore{}
}

func (s *LunchStore) Create(l slash_commands.Lunch) (*slash_commands.Lunch, error) {
	r := &slash_commands.Lunch{}
	fmt.Println("The lunch was created successfully.")
	return r, nil
}
