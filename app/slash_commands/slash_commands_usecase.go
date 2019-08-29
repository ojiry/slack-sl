package slash_commands

type LunchRepository interface {
	Create(l Lunch) (*Lunch, error)
}

type SlashCommandsUsecase struct {
	repo LunchRepository
}

func NewSlashCommandsUsecase(repo LunchRepository) *SlashCommandsUsecase {
	return &SlashCommandsUsecase{repo: repo}
}

func (u *SlashCommandsUsecase) CreateLunch(sc SlashCommands) (*Lunch, error) {
	l, err := u.repo.Create(Lunch{})
	if err != nil {
		return nil, err
	}

	return l, nil
}
