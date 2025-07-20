package repository

import (
	"database/sql"
	"github.com/DiamondDmitriy/big-note-api/internal/core/entity/task"
)

type TaskRepository struct {
	//db *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{}
}

var tasksData = []task.Task{
	{
		Id:         1,
		Order:      1,
		Name:       "Нужно сделать кайфовую API",
		Done:       false,
		CategoryId: 1,
	},
	{
		Id:         2,
		Order:      3,
		Name:       "Нужно сделать",
		Done:       true,
		CategoryId: 2,
	},
	{
		Id:         3,
		Order:      2,
		Name:       "Нужно сделать",
		Done:       false,
		CategoryId: 3,
	},
}

func (r *TaskRepository) GetAll() ([]task.Task, error) {
	return tasksData, nil
}

func (r *TaskRepository) Get(id int) (task.Task, error) {
	return tasksData[0], nil
}

func (r *TaskRepository) GetAllGroupBy(task task.Task) {
}

//func (t * TaskRepository) Insert(task types.Task) (types.Task, error) {}
