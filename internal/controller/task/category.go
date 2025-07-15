package task

import (
	"database/sql"
	"errors"
	"github.com/DiamondDmitriy/big-note-api/internal/repository"
	"github.com/DiamondDmitriy/big-note-api/internal/service"
	"github.com/DiamondDmitriy/big-note-api/pkg/rest"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CategoryController struct {
	Repo     *repository.TaskCategoryRepository
	TaskRepo *repository.TaskRepository
	Service  *service.TaskCategoryService
}

func (c *CategoryController) TaskCategoryGetOne(ctx *gin.Context) {
	var requestData struct {
		Id int `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&requestData); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	category, err := c.Repo.GetOne(requestData.Id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rest.ResponseError(ctx, http.StatusNotFound, "Not found", err.Error())
			return
		}

		rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
		return
	}

	rest.ResponseSuccess(ctx, http.StatusOK, category, nil)
}

func (c *CategoryController) TaskCategoryGetAll(ctx *gin.Context) {
	includeTasks := ctx.Query("include_tasks") == "true"

	categories, err := c.Repo.GetAll()
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
		return
	}

	if !includeTasks {
		rest.ResponseSuccess(ctx, http.StatusOK, categories, nil)
		return
	}

	// Вместе с тасками
	//categoryService := service.NewTaskCategoryService(c.TaskRepo)
	//categoriesWithTasks, err := categoryService.GetCategoriesWithTasks(categories)
	//if err != nil {
	//	rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
	//	return
	//}

	//rest.ResponseSuccess(ctx, http.StatusOK, categoriesWithTasks, nil)
}

// TaskCategoryCreate Create
func (c *CategoryController) TaskCategoryCreate(ctx *gin.Context) {
	var requestData struct {
		Name string `json:"name" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	if requestData.Name == "" {
		rest.ResponseError(ctx, http.StatusBadRequest, "Name cannot be empty", nil)
		return
	}

	category, err := c.Repo.Create(requestData.Name)

	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, "Failed to create category", err.Error())
		return
	}

	rest.ResponseSuccess(ctx, http.StatusCreated, category, nil)
}

// TaskCategoryDelete Delete
func (c *CategoryController) TaskCategoryDelete(ctx *gin.Context) {
	var requestData struct {
		Id int `uri:"id" binding:"required"`
	}

	if err := ctx.ShouldBindUri(&requestData); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	if err := c.Repo.Delete(requestData.Id); err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
		return
	}

	rest.ResponseSuccess(ctx, http.StatusOK, requestData, nil)
}

// TaskCategoryUpdate Update
func (c *CategoryController) TaskCategoryUpdate(ctx *gin.Context) {
	var requestData struct {
		Name string `json:"name" binding:"required"`
	}

	idRaw := ctx.Param("id")
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "id is empty", err.Error())
		return
	}

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	err = c.Repo.Update(id, requestData.Name)
	if err != nil {
		return
	}

	rest.ResponseSuccess(ctx, http.StatusOK, requestData, nil)
}
