package service

// import (
//
//	"fmt"
//	"github.com/golang-jwt/jwt/v5"
//	"github.com/spf13/viper"
//	"log"
//
// )
type AuthService struct{}

//func (a *Auth) JWTVerifyUser(token string) bool {
//	//ParseWithClaims
//	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
//		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
//		}
//
//		tokenPassword := viper.GetString("token_password")
//		return []byte(tokenPassword), nil
//	})
//	if err != nil {
//		log.Fatalf("Failed to parse token: %s\n", err)
//		return false
//	}
//
//	if !jwtToken.Valid {
//		return false
//	}
//
//	claims, ok := jwtToken.Claims.(jwt.MapClaims)
//	if !ok {
//		return false
//	}
//
//	loginRaw, ok := claims["login"]
//	if !ok {
//		return false
//	}
//
//	_, ok = loginRaw.(string)
//	if !ok {
//		return false
//	}
//
//	// todo: Проверка пользователя
//	return true
//
//	//for _, roles := range userRoles[login] {
//	//	for _, storedPermission := range rolePermissions[roles] {
//	//		if permission == storedPermission {
//	//			return true
//	//		}
//	//	}
//	//}
//	return false
//}
//
//func validate(login, password string) bool {
//	return true
//}
//
//func (a *Auth) Create() {}
//func (a *Auth) Update() {}
//func (a *Auth) Delete() {}
//func (a *Auth) Login()  {}
