package repository

import (
	"database/sql"
	entity "github.com/DiamondDmitriy/big-note-api/internal/core/entity/task"
)

type TaskCategoryRepository struct {
	db *sql.DB
}

func NewTaskCategoryRepository(db *sql.DB) *TaskCategoryRepository {
	return &TaskCategoryRepository{db}
}

func (r *TaskCategoryRepository) scanTaskCategory(row *sql.Row) (entity.Category, error) {
	var taskCategory entity.Category
	err := row.Scan(
		&taskCategory.Id,
		&taskCategory.Name,
		&taskCategory.CreatedAt,
		&taskCategory.UpdatedAt,
		&taskCategory.UserId,
	)
	return taskCategory, err
}

func (r *TaskCategoryRepository) GetAll(userId string) ([]entity.Category, error) {
	var taskCategories []entity.Category
	query := "SELECT id ,name, created_at, updated_at, user_id FROM todo.categories WHERE user_id=$1"
	rows, err := r.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var taskCategory entity.Category
		// todo: повторяется
		if err := rows.Scan(
			&taskCategory.Id,
			&taskCategory.Name,
			&taskCategory.CreatedAt,
			&taskCategory.UpdatedAt,
			&taskCategory.UserId,
		); err != nil {
			return nil, err
		}

		taskCategories = append(taskCategories, taskCategory)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return taskCategories, nil
}

func (r *TaskCategoryRepository) GetOne(id int) (entity.Category, error) {
	query := "SELECT id ,name, created_at, updated_at FROM todo.categories WHERE id=$1"
	row := r.db.QueryRow(query, id)
	category, err := r.scanTaskCategory(row)
	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (r *TaskCategoryRepository) Create(name string, userId string) (entity.Category, error) {
	query := "INSERT INTO todo.categories (name,user_id) VALUES($1, $2) RETURNING id, name, created_at, updated_at, user_id"
	row := r.db.QueryRow(query, name, userId)
	category, err := r.scanTaskCategory(row)

	if err != nil {
		return entity.Category{}, err
	}
	return category, nil
}

func (r *TaskCategoryRepository) Update(id int, name string) error {
	return nil
}

func (r *TaskCategoryRepository) Delete(id int) error {
	query := "DELETE FROM todo.categories WHERE id=$1"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
