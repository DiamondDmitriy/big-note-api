package repository

import "database/sql"

type Repository struct {
	Auth         *AuthRepository
	Task         *TaskRepository
	TaskCategory *TaskCategoryRepository
}

func NewRepositories(db *sql.DB) *Repository {
	return &Repository{
		Auth:         NewAuthRepository(db),
		Task:         NewTaskRepository(db),
		TaskCategory: NewTaskCategoryRepository(db),
	}
}
