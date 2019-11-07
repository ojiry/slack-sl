package slash_commands

type LunchRepository interface {
	Create(lunch Lunch) (*Lunch, error)
}

type Lunch struct {
	ID    uint
	Users []User
}
