package service

import (
	"github.com/DiamondDmitriy/big-note-api/internal/entity/task"
	"github.com/DiamondDmitriy/big-note-api/internal/repository"
)

type TaskCategoryService struct {
	TaskRepository *repository.TaskRepository
}

func NewTaskCategoryService(tr *repository.TaskRepository) *TaskCategoryService {
	return &TaskCategoryService{tr}
}

// GetCategoriesWithTasks Вернёт категории вместе с тасками
func (s *TaskCategoryService) GetCategoriesWithTasks(categories []task.Category) ([]task.TasksByCategory, error) {
	var respData []task.TasksByCategory

	// Задачи
	tasks, err := s.TaskRepository.GetAll()
	if err != nil {
		return nil, err
	}

	tasksByCategory := make(map[int][]task.Task)
	for _, taskEl := range tasks {
		tasksByCategory[taskEl.CategoryId] = append(tasksByCategory[taskEl.CategoryId], taskEl)
	}

	for _, category := range categories {
		respData = append(respData, task.TasksByCategory{
			Category: category,
			Tasks:    tasksByCategory[category.Id],
		})
	}

	return respData, nil
}
