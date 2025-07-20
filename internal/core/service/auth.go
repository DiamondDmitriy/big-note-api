package service

import (
	"github.com/DiamondDmitriy/big-note-api/config"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService struct {
	Config *config.Config
	//Repo     *auth.Repository
	//UserRepo *user.Repository
}

type Claims struct {
	jwt.RegisteredClaims
}

//// CreateTokenJWT
//// 5. Что будет, если проигнорировать aud?
//// Уязвимость к атакам: Злоумышленник может использовать токен, украденный у Сервиса A, в Сервисе B.
////
//// Нарушение принципа минимальных привилегий: Токен будет работать везде, где принимается.
////
//// 6. Дополнительные рекомендации
//// Всегда проверяйте aud на стороне сервера.
////
//// Используйте разные ключи для разных аудиторий (если требуется максимальная изоляция).
////
//// Включайте aud даже если у вас один сервис — это улучшит безопасность.
//func (a *AuthService) CreateTokenJWT(user *userentity.User) (string, error) {
//	claims := &Claims{
//		RegisteredClaims: jwt.RegisteredClaims{
//			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Срок действия - 24 часа
//			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // Время выдачи
//			NotBefore: jwt.NewNumericDate(time.Now()),                     // Действителен с
//			Issuer:    a.Config.APP.Name,                                  // Кто выдал
//			Subject:   user.Username,                                      // Для кого
//			ID:        user.Id,                                            // Уникальный ID токена
//			Audience:  []string{"big_note_api_todo"},                      // Аудитория - для каких сервисов доступен токен
//		},
//	}
//
//	// Создаем токен с методом подписи HS256
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	// Подписываем токен секретным ключом
//	signedToken, err := token.SignedString(a.Config.JWT.TokenPassword)
//	if err != nil {
//		return signedToken, err
//	}
//
//	return signedToken, nil
//}
//
////// RevokeToken Отозвать токен
////func (a *AuthService) RevokeToken() {}
////
////// VerifyUserTokenJWT Проверка access токена
////func VerifyUserTokenJWT(token string, tokenPassword []byte) (bool, *Claims) {
////	claims := &Claims{}
////	jwtToken, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
////		return tokenPassword, nil
////	})
////
////	if err != nil {
////		return false, claims
////	}
////
////	if !jwtToken.Valid {
////		fmt.Println("ne valid")
////		return false, claims
////	}
////
////	// todo: Проверка пользователя
////	return true, claims
////
////	////for _, roles := range userRoles[login] {
////	////	for _, storedPermission := range rolePermissions[roles] {
////	////		if permission == storedPermission {
////	////			return true
////	////		}
////	////	}
////	////}
////}
////
////// Authenticate Авторизация
////func (a *AuthService) Authenticate(loginOrEmail string, password string) (*user.User, error) {
////	// ищем пользователя в бд
////	user, err := a.UserRepo.FindByLoginOrEmail(loginOrEmail)
////	if err != nil {
////		return user, err
////	}
////
////	// Проверка пароля
////	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
////	if err != nil {
////		return new(user.User), err
////	}
////
////	return user, nil
////}
//
//// Registration Регистрация
//func (a *AuthService) Registration(usr *userentity.Registration) (*userentity.User, error) {
//	// Генерация хеша (стоимость = 10, можно увеличить для большей безопасности)
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(usr.Password), bcrypt.DefaultCost)
//	if err != nil {
//		panic(err)
//	}
//
//	usr.PasswordHash = hashedPassword
//
//	newUsr, err := a.Repo.Registration(usr)
//	if err != nil {
//		return newUsr, err
//	}
//
//	//if err != nil {
//	//	var pqErr *pq.Error
//	//	if errors.As(err, &pqErr) {
//	//		if pqErr.Code == "23505" && pqErr.Constraint == "users_unique_username" {
//	//			rest.ResponseError(ctx, http.StatusBadRequest, "Этот логин уже занят", err.Error())
//	//			return
//	//		} else if pqErr.Code == "23505" && pqErr.Constraint == "users_unique_email" {
//	//			rest.ResponseError(ctx, http.StatusBadRequest, "Эта почта уже занят", err.Error())
//	//			return
//	//		}
//	//		fmt.Println(pqErr.Code, pqErr.Message)
//	//	}
//	//
//	//	rest.ResponseError(ctx, http.StatusInternalServerError, "Server operation failed", err.Error())
//	//	return
//	//}
//
//	return newUsr, nil
//}
