package controller

import (
	"github.com/DiamondDmitriy/big-note-api/internal/controller/auth"
	"github.com/DiamondDmitriy/big-note-api/internal/controller/task"
	"github.com/DiamondDmitriy/big-note-api/internal/repository"
	"github.com/DiamondDmitriy/big-note-api/internal/service"
)

type Controller struct {
	Task         *task.Controller
	TaskCategory *task.CategoryController
	Auth         *auth.Controller
}

func NewControllers(repositories *repository.Repository, services *service.Service) *Controller {
	return &Controller{
		Task: &task.Controller{
			Repository: repositories.Task,
		},
		TaskCategory: &task.CategoryController{
			Repo:     repositories.TaskCategory,
			TaskRepo: repositories.Task,
			Service:  services.TaskCategory,
		},
		Auth: &auth.Controller{
			Repo:    repositories.Auth,
			Service: services.Auth,
		},
	}
}
