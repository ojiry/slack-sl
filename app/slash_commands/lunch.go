package slash_commands

type LunchRepository interface {
	Find(lunchID int64) *Lunch
	Create(lunch Lunch) (*Lunch, error)
}

type Lunch struct {
	ID    int64
	Users []User
}
