package service

import (
	"github.com/DiamondDmitriy/big-note-api/config"
	"github.com/DiamondDmitriy/big-note-api/internal/repository"
)

type Service struct {
	TaskCategory *TaskCategoryService
	Auth         *AuthService
}

func NewServices(repos *repository.Repository, cnf *config.Config) *Service {
	return &Service{
		TaskCategory: NewTaskCategoryService(repos.Task),
		Auth: &AuthService{
			cnf,
			repos.Auth,
			repos.User,
		},
	}
}
