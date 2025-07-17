package user

import (
	"database/sql"
	entity "github.com/DiamondDmitriy/big-note-api/internal/entity/user"
)

type Repository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) ScanUser(row *sql.Row) (entity.User, error) {
	var user entity.User
	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Name,
		&user.Surname,
		//&user.Ro
	)
	return user, err
}

func (r *Repository) findById(id int) (*entity.User, error) {
	query := "SELECT id username email password_hash created_at updated_at name surname role_id FROM USERS.USERS WHERE ID = $1"
	row := r.db.QueryRow(query, id)
	user, err := r.ScanUser(row)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) FindByLoginOrEmail(loginOrMail string) (*entity.User, error) {
	//todo: role
	query := `
SELECT id, username, email, password_hash, created_at, updated_at, name, surname
FROM users.users WHERE username = $1 or email = $1
`
	row := r.db.QueryRow(query, loginOrMail)
	user, err := r.ScanUser(row)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
