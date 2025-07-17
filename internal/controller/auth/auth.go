package auth

import (
	"database/sql"
	"errors"
	"fmt"
	authEntity "github.com/DiamondDmitriy/big-note-api/internal/entity/auth"
	userEntity "github.com/DiamondDmitriy/big-note-api/internal/entity/user"
	"github.com/DiamondDmitriy/big-note-api/internal/repository/auth"
	"github.com/DiamondDmitriy/big-note-api/internal/service"
	"github.com/DiamondDmitriy/big-note-api/pkg/rest"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type Controller struct {
	Repo    *auth.Repository
	Service *service.AuthService
}

// SignUp Регистрация
func (c *Controller) SignUp(ctx *gin.Context) {
	var requestData userEntity.Registration

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	user, err := c.Service.Registration(&requestData)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == "23505" && pqErr.Constraint == "users_unique_username" {
				rest.ResponseError(ctx, http.StatusBadRequest, "Этот логин уже занят", err.Error())
				return
			} else if pqErr.Code == "23505" && pqErr.Constraint == "users_unique_email" {
				rest.ResponseError(ctx, http.StatusBadRequest, "Эта почта уже занят", err.Error())
				return
			}
			fmt.Println(pqErr.Code, pqErr.Message)
		}

		rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
		return
	}

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
func (c *Controller) SignIn(ctx *gin.Context) {
	var requestData struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestData); err != nil {
		rest.ResponseError(ctx, http.StatusBadRequest, "Invalid request format", err.Error())
		return
	}

	user, err := c.Service.Authenticate(requestData.Login, requestData.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			rest.ResponseError(ctx, http.StatusUnauthorized, "Пользователь не найден", err.Error())
			return
		}

		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			rest.ResponseError(ctx, http.StatusUnauthorized, "Неверный пароль", err.Error())
			return
		}

		rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
		return
	}

	user.PasswordHash = nil
	jwtToken, err := c.Service.CreateTokenJWT(user)

	rest.ResponseSuccess(ctx, http.StatusOK, authEntity.UserWithToken{
		Token: jwtToken,
		User:  user,
	}, nil)
}
