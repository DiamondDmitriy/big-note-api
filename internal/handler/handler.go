package handler

import (
	"github.com/gin-gonic/gin"
)

//import (
//	"errors"
//	"github.com/DiamondDmitriy/gopher-todo/internal/service"
//	"github.com/gin-gonic/gin"
//	"net/http"
//	"strings"
//)

type Handler struct{}

//func authenticationUser(ctx *gin.Context) {
//	authHeaderValue := ctx.GetHeader("Authorization")
//
//	if authHeaderValue != "" {
//		bearerToken := strings.Split(authHeaderValue, " ")
//		if len(bearerToken) == 2 && bearerToken[0] == "Bearer" {
//			authService := new(service.Auth)
//
//			if authService.JWTVerifyUser(bearerToken[1]) {
//				ctx.Next()
//			}
//		}
//	}
//
//	ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
//}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.sighIn)
		auth.POST("/sign-up", h.sighUp)
	}

	// //router.ContextWithFallback

	//api := router.Group("/api")
	//api.Use(authenticationUser)
	//api.GET("/tasks", h.tasksGetAll)
	//
	//task := api.Group("/task")
	//{
	//	task.GET("/:id", h.taskGetById)
	//	task.POST("/:id", h.taskAdd)
	//	task.PATCH("/:id", h.taskEdit)
	//	task.DELETE("/:id", h.taskDeleteById)
	//}

	return router
}
