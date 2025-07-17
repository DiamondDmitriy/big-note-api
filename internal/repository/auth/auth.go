package auth

import (
	"database/sql"
	"github.com/DiamondDmitriy/big-note-api/internal/entity/user"
	userRepo "github.com/DiamondDmitriy/big-note-api/internal/repository/user"
)

type Repository struct {
	db       *sql.DB
	UserRepo *userRepo.Repository
}

func NewAuthRepository(db *sql.DB, repo *userRepo.Repository) *Repository {
	return &Repository{db, repo}
}

func (r *Repository) Registration(usr *user.Registration) (*user.User, error) {
	query := `INSERT INTO users.users
(username, email, password_hash, "name", surname, role_id)
VALUES($1, $2, $3,$4, $5, $6) RETURNING id, username, email, password_hash, created_at, updated_at, "name", surname;
`
	row := r.db.QueryRow(query, usr.Login, usr.Email, usr.PasswordHash, usr.Name, usr.Surname, nil)

	newUser, err := r.UserRepo.ScanUser(row)
	if err != nil {
		return &newUser, err
	}

	return &newUser, nil
}

//func userScan(row *sql.Row) *user.User {
//	usr := &user.User{}
//	row.Scan(&usr.Id)
//}

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
