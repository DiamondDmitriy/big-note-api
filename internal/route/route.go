package route

import (
	"github.com/DiamondDmitriy/big-note-api/config"
	"github.com/DiamondDmitriy/big-note-api/internal/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

type Route struct {
	Config     *config.Config
	Controller *controller.Controller
}

func authenticationUser(ctx *gin.Context) {
	//authHeaderValue := ctx.GetHeader("Authorization")
	//
	//if authHeaderValue != "" {
	//	bearerToken := strings.Split(authHeaderValue, " ")
	//	if len(bearerToken) == 2 && bearerToken[0] == "Bearer" {
	//		authService := new(service.Auth)
	//
	//		if authService.JWTVerifyUser(bearerToken[1]) {
	//			ctx.Next()
	//		}
	//	}
	//}
	//
	//ctx.AbortWithError(http.StatusUnauthorized, errors.New("unauthorized"))
}

// Конфигурация CORS
func (r *Route) getCors() gin.HandlerFunc {
	// Получаем список разрешённых доменов из переменных окружения
	allowOrigins := strings.Split(r.Config.HTTP.CorsAllowOrigins, ",")

	// Удаляем возможные пробелы вокруг доменов
	for i, origin := range allowOrigins {
		allowOrigins[i] = strings.TrimSpace(origin)
	}

	return cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodPatch, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length", "X-Total-Count"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour, // Кеширование предварительных запросов
	})
}

func (r *Route) Init() *gin.Engine {
	router := gin.New()
	router.Use(
		gin.Recovery(), // Восстановление после паник
		r.getCors(),    // CORS
		gin.Logger(),   // Логирование
	)

	//r.AuthRoute(router)

	//// Public routes
	//auth := router.Group("/auth")
	//{
	//	auth.POST("/sign-in", r.Handlers.sighIn)
	//	auth.POST("/sign-up", h.sighUp)
	//}
	//
	//// //router.ContextWithFallback
	//
	api := router.Group("/api")
	// Protected API routes
	api.Use(authenticationUser)
	r.TaskRoutes(api)

	return router
}
