package repository

import (
	"github.com/ojiry/slack-sl/app/domain"
)

type LunchRepository struct{}

func (r *LunchRepository) Create(*domain.User) (*domain.Lunch, error) {
	// TODO: Implement the persistence process
	return nil, nil
}
