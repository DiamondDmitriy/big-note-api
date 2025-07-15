package route

import (
	"github.com/gin-gonic/gin"
)

func (r *Route) TaskRoutes(api *gin.RouterGroup) {
	ctrlTask := r.Controller.Task
	ctrlCategory := r.Controller.TaskCategory

	tasks := api.Group("/tasks")
	{
		tasks.GET("", ctrlTask.TasksGetAll)
		tasks.GET("/categories", ctrlCategory.TaskCategoryGetAll)
	}

	task := api.Group("/task")
	{
		task.GET("/:id", ctrlTask.TaskGetOne)
		task.POST("", ctrlTask.TaskCreate)
		task.PUT("/:id", ctrlTask.TaskUpdate) // PUT для полного обновления
		//task.PATCH("/:id", nil)         // PATCH для частичного
		task.DELETE("/:id", ctrlTask.TaskDelete)
	}

	taskCategory := tasks.Group("/category")
	{
		taskCategory.GET("/:id", ctrlCategory.TaskCategoryGetOne)
		taskCategory.POST("", ctrlCategory.TaskCategoryCreate)
		taskCategory.PATCH("/:id", ctrlCategory.TaskCategoryUpdate)
		taskCategory.DELETE("/:id", ctrlCategory.TaskCategoryDelete)
	}
}
