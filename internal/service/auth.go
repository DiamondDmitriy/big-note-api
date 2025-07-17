package service

import (
	"fmt"
	"github.com/DiamondDmitriy/big-note-api/config"
	userEntity "github.com/DiamondDmitriy/big-note-api/internal/entity/user"
	"github.com/DiamondDmitriy/big-note-api/internal/repository/auth"
	userRepo "github.com/DiamondDmitriy/big-note-api/internal/repository/user"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthService struct {
	Config   *config.Config
	Repo     *auth.Repository
	UserRepo *userRepo.Repository
}

type Claims struct {
	jwt.RegisteredClaims
}

// CreateTokenJWT
// 5. Что будет, если проигнорировать aud?
// Уязвимость к атакам: Злоумышленник может использовать токен, украденный у Сервиса A, в Сервисе B.
//
// Нарушение принципа минимальных привилегий: Токен будет работать везде, где принимается.
//
// 6. Дополнительные рекомендации
// Всегда проверяйте aud на стороне сервера.
//
// Используйте разные ключи для разных аудиторий (если требуется максимальная изоляция).
//
// Включайте aud даже если у вас один сервис — это улучшит безопасность.
func (a *AuthService) CreateTokenJWT(user *userEntity.User) (string, error) {
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Срок действия - 24 часа
		IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Время выдачи
		NotBefore: jwt.NewNumericDate(time.Now()),                     // Действителен с
		Issuer:    a.Config.APP.Name,                                  // Кто выдал
		Subject:   user.Username,                                      // Для кого
		ID:        user.Id,                                            // Уникальный ID токена
		Audience:  []string{"big_note_api_todo"},                      // Аудитория - для каких сервисов доступен токен
	}

	// Создаем токен с методом подписи HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	signedToken, err := token.SignedString(a.Config.JWT.TokenPassword)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}

// RevokeToken Отозвать токен
func (a *AuthService) RevokeToken() {}

// VerifyUserTokenJWT Проверка access токена
func VerifyUserTokenJWT(token string, tokenPassword []byte) bool {
	claims := &Claims{}
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return tokenPassword, nil
	})

	if err != nil {
		fmt.Println(err)
		//fmt.Println()
		return false
	}
	fmt.Println(claims.Subject)
	fmt.Println()

	if !jwtToken.Valid {
		fmt.Println("ne valid")
		return false
	}

	// todo: Проверка пользователя
	return true

	////for _, roles := range userRoles[login] {
	////	for _, storedPermission := range rolePermissions[roles] {
	////		if permission == storedPermission {
	////			return true
	////		}
	////	}
	////}
}

// Authenticate Авторизация
func (a *AuthService) Authenticate(loginOrEmail string, password string) (*userEntity.User, error) {
	// ищем пользователя в бд
	user, err := a.UserRepo.FindByLoginOrEmail(loginOrEmail)
	if err != nil {
		return user, err
	}

	// Проверка пароля
	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		return new(userEntity.User), err
	}

	return user, nil
}

// Registration Регистрация
func (a *AuthService) Registration(usr *userEntity.Registration) (*userEntity.User, error) {
	// Генерация хеша (стоимость = 10, можно увеличить для большей безопасности)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	usr.PasswordHash = hashedPassword

	newUsr, err := a.Repo.Registration(usr)
	if err != nil {
		return newUsr, err
	}

	return newUsr, nil
}
