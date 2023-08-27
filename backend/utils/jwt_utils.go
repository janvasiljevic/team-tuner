package utils

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type UserRole string

const (
	UserRoleAdmin   UserRole = "admin"
	UserRoleStudent UserRole = "student"
)

var UserRoleLookup = map[string]UserRole{
	"admin":   UserRoleAdmin,
	"student": UserRoleStudent,
}

var JWTSecret = []byte("!!SECRET!!")

func GenerateJWTCookie(userId uuid.UUID, role UserRole) *http.Cookie {
	return &http.Cookie{
		Name:     "jwt",
		Value:    generateJWT(userId, role),
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	}
}

func ClearJWtCookie() *http.Cookie {
	return &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		Expires:  time.Now().Add(-24 * time.Hour),
		HttpOnly: true,
	}
}

func generateJWT(id uuid.UUID, role UserRole) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["id"] = id
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, _ := token.SignedString(JWTSecret)

	return t
}
