package authhandle

import (
	"github.com/DiamondDmitriy/big-note-api/internal/core/entity/user"
	"github.com/DiamondDmitriy/big-note-api/internal/core/service"
	"github.com/DiamondDmitriy/big-note-api/pkg/rest"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service *service.AuthService
}

// SignUp Регистрация
func (c *Handler) SignUp(ctx *gin.Context) {
	var requestData userentity.Registration

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	user, err := c.Service.Registration(&requestData)

	user.PasswordHash = nil
	jwtToken, err := c.Service.CreateTokenJWT(user)
	if err != nil {
		rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
		return
	}

	rest.ResponseSuccess(ctx, http.StatusOK, authEntity.UserWithToken{
		Token: jwtToken,
		User:  user,
	}, nil)
}

// SignIn Вход
func (c *Handler) SignIn(ctx *gin.Context) {
	//	var requestData struct {
	//		Login    string `json:"login" binding:"required"`
	//		Password string `json:"password" binding:"required"`
	//	}
	//
	//	if err := ctx.ShouldBindJSON(&requestData); err != nil {
	//		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
	//		return
	//	}
	//
	//	user, err := c.Service.Authenticate(requestData.Login, requestData.Password)
	//	if err != nil {
	//		if errors.Is(err, sql.ErrNoRows) {
	//			rest.ResponseError(ctx, http.StatusUnauthorized, "Пользователь не найден", err.Error())
	//			return
	//		}
	//
	//		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
	//			rest.ResponseError(ctx, http.StatusUnauthorized, "Неверный пароль", err.Error())
	//			return
	//		}
	//
	//		rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
	//		return
	//	}
	//
	//	user.PasswordHash = nil
	//	jwtToken, err := c.Service.CreateTokenJWT(user)
	//
	//	rest.ResponseSuccess(ctx, http.StatusOK, authEntity.UserWithToken{
	//		Token: jwtToken,
	//		User:  user,
	//	}, nil)
}
