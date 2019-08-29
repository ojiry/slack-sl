package usecase

import (
	"github.com/ojiry/slack-sl/app/domain"
)

type LunchRepository interface {
	Create(*domain.User) (*domain.Lunch, error)
}

type CreateLunchUsecase struct {
	repo LunchRepository
}

func NewCreateLunchUsecase(repo LunchRepository) *CreateLunchUsecase {
	return &CreateLunchUsecase{repo: repo}
}

func (u *CreateLunchUsecase) Create(user *domain.User) (*domain.Lunch, error) {
	l, err := u.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return l, nil
}
