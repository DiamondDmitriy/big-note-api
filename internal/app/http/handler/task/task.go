package task

import (
	"github.com/DiamondDmitriy/big-note-api/internal/repository"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Repository *repository.TaskRepository
}

func (c *Controller) TasksGetAll(ctx *gin.Context) {
	//tasks := taskRepository.GetAll()
	//ctx.IndentedJSON(200, tasks)
}

func (c *Controller) TaskGetOne(ctx *gin.Context) {}

func (c *Controller) TaskCreate(ctx *gin.Context) {}

func (c *Controller) TaskDelete(ctx *gin.Context) {}

func (c *Controller) TaskUpdate(ctx *gin.Context) {}
