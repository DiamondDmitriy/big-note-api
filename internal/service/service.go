package service

import "github.com/DiamondDmitriy/big-note-api/internal/repository"

type Service struct {
	TaskCategory *TaskCategoryService
	//Auth         *AuthService
}

func NewServices(repos *repository.Repository) *Service {
	return &Service{
		TaskCategory: NewTaskCategoryService(repos.Task),
	}
}
