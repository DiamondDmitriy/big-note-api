package auth

import userEntity "github.com/DiamondDmitriy/big-note-api/internal/entity/user"

type UserWithToken struct {
	Token string           `json:"token"`
	User  *userEntity.User `json:"user"`
}
