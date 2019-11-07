package app

import (
	"errors"
	"net/http"
	"os"

	scHandler "github.com/ojiry/slack-sl/app/slash_commands/handler"
	scStore "github.com/ojiry/slack-sl/app/slash_commands/store"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		return errors.New("$PORT must be set")
	}

	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		return errors.New("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&scStore.Lunch{})

	http.HandleFunc("/slash_commands", scHandler.New(db))

	return http.ListenAndServe(":"+port, nil)
}
