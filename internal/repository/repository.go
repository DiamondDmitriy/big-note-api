package repository

import (
	"database/sql"
	"github.com/DiamondDmitriy/big-note-api/internal/repository/auth"
	"github.com/DiamondDmitriy/big-note-api/internal/repository/user"
)

type Repository struct {
	Auth         *auth.Repository
	Task         *TaskRepository
	TaskCategory *TaskCategoryRepository
	User         *user.Repository
}

func NewRepositories(db *sql.DB) *Repository {
	userRepo := user.NewUserRepository(db)
	return &Repository{
		Auth:         auth.NewAuthRepository(db, userRepo),
		Task:         NewTaskRepository(db),
		TaskCategory: NewTaskCategoryRepository(db),
		User:         userRepo,
	}
}
