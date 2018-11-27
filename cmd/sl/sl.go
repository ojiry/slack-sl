package main

import (
	"net/http"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
)

type Channel struct {
	gorm.Model
	SlackID string
	Name    string
	Lunch   []Lunch
}

type Group struct {
	gorm.Model
	Lunch   Lunch
	LunchID uint
	Members []User `gorm:"many2many:group_members;"`
}

type User struct {
	gorm.Model
	SlackID  string
	Username string
}

type Lunch struct {
	gorm.Model
	Channel      Channel
	ChannelID    uint
	Groups       []Group
	ShuffledAt   *time.Time
	Participants []User `gorm:"many2many:participations;"`
}

type InteractiveMessage struct {
	Text         string `json:"text"`
	ResponseType string `json:"response_type"`
	Attachments  []struct {
		Text           string `json:"text"`
		Fallback       string `json:"fallback"`
		CallbackID     string `json:"callback_id"`
		Color          string `json:"color"`
		AttachmentType string `json:"attachment_type"`
		Actions        []struct {
			Name    string `json:"name"`
			Text    string `json:"text"`
			Type    string `json:"type"`
			Value   string `json:"value"`
			Style   string `json:"style,omitempty"`
			Confirm struct {
				Title       string `json:"title"`
				Text        string `json:"text"`
				OkText      string `json:"ok_text"`
				DismissText string `json:"dismiss_text"`
			} `json:"confirm,omitempty"`
		} `json:"actions"`
	} `json:"attachments"`
}

func main() {
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	db.AutoMigrate(&Channel{}, &User{}, &Lunch{}, &Group{})

	e := echo.New()
	e.POST("/slash_commands", slashCommandsHandler)
	e.GET("/slash_commands", slashCommandsHandler)
	e.Start(":8080")
}

func slashCommandsHandler(c echo.Context) error {
	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()
	var channel Channel
	db.FirstOrCreate(&channel, Channel{SlackID: c.Param("channel_id"), Name: c.Param("channel_name")})
	var lunch Lunch
	db.Where(Lunch{ChannelID: channel.ID, ShuffledAt: nil}).FirstOrCreate(&lunch)
	return c.JSON(http.StatusOK, InteractiveMessage{
		Text:         "Would you like to join the Shuffle Lunch today?",
		ResponseType: "in_channel",
	})
}
