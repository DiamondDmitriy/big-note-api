package auth

import (
	userEntity "github.com/DiamondDmitriy/big-note-api/internal/core/entity/user"
)

type UserWithToken struct {
	Token string           `json:"token"`
	User  *userEntity.User `json:"user"`
}
