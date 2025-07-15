package user

type User struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Role       string `json:"role"`
}

// Login todo: авторизация по логину или почте
type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
