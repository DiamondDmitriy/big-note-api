package user

type Registration struct {
	Login        string `json:"login" binding:"required"`
	Password     string `json:"password" binding:"required"`
	PasswordHash []byte `json:"password_hash"`
	Email        string `json:"email" binding:"required"`
	Name         string `json:"name" binding:"required"`
	Surname      string `json:"surname" binding:"required"`
}
