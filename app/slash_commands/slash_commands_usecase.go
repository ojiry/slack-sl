package slash_commands

type SlashCommandsUsecase struct {
	repo LunchRepository
}

func NewSlashCommandsUsecase(repo LunchRepository) *SlashCommandsUsecase {
	return &SlashCommandsUsecase{repo: repo}
}

func (u *SlashCommandsUsecase) CreateLunch(sc *SlashCommands) (*Lunch, error) {
	lunch, err := u.repo.Create(Lunch{})
	if err != nil {
		return nil, err
	}

	return lunch, nil
}
