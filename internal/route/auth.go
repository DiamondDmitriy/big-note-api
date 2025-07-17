package route

import "github.com/gin-gonic/gin"

func (r *Route) AuthRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", r.Controller.Auth.SignIn)
		auth.POST("/sign-up", r.Controller.Auth.SignUp)
	}
}
