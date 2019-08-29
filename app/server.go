package app

import (
	"github.com/ojiry/slack-sl/app/repository"
	"github.com/ojiry/slack-sl/app/usecase"
)

type Server struct {
	usecase *usecase.CreateLunchUsecase
}

func NewServer() *Server {
	repo := &repository.LunchRepository{}
	return &Server{
		usecase: usecase.NewCreateLunchUsecase(repo),
	}
}
