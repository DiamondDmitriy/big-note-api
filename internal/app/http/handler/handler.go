package handler

import (
	"github.com/DiamondDmitriy/big-note-api/internal/app/http/handler/auth"
	"github.com/DiamondDmitriy/big-note-api/internal/core/service"
)

//import (
//	"github.com/DiamondDmitriy/big-note-api/internal/core/service"
//	//"github.com/DiamondDmitriy/big-note-api/internal/repository"
//	"github.com/DiamondDmitriy/big-note-api/internal/app/http/handler/auth/authcontroller"
//)

type Handler struct {
	Auth *authhandle.Handler
	//Task         *task2.Controller
	//TaskCategory *task2.CategoryController
}

func New(services *service.Service) *Handler {
	return &Handler{
		//Task: &task2.Controller{
		//	Repository: repositories.Task,
		//},
		//TaskCategory: &task2.CategoryController{
		//	Repo:     repositories.TaskCategory,
		//	TaskRepo: repositories.Task,
		//	Service:  services.TaskCategory,
		//},
		Auth: &authhandle.Handler{
			Service: services.Auth,
		},
	}
}
