package repository

import "database/sql"

type AuthRepository struct{}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{}
}

//func (a *Auth) SighIn(username, password string) bool {
//	return true
//}
//
//func (a *Auth) SighUp(username, password string) bool {
//	return true
//}
//
//func (a *Auth) SighOut(username, password string) {}
//
//func (a *Auth) GetUser() {}
