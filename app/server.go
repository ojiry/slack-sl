package app

import (
	"slack-sl/app/repository"
	"slack-sl/app/usecase"
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
