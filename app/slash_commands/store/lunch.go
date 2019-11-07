package store

import (
	"fmt"

	"github.com/ojiry/slack-sl/app/slash_commands"

	"github.com/jinzhu/gorm"
)

type Lunch struct {
	gorm.Model
}

type LunchStore struct {
	db *gorm.DB
}

func NewLunchStore(db *gorm.DB) *LunchStore {
	return &LunchStore{db: db}
}

func (s *LunchStore) Create(l slash_commands.Lunch) (*slash_commands.Lunch, error) {
	lunch := Lunch{}
	s.db.Create(&lunch)
	fmt.Printf("The lunch was created successfully. %v", lunch)
	r := &slash_commands.Lunch{ID: lunch.ID}
	return r, nil
}
