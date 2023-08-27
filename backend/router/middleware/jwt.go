package middleware

import (
	"fmt"
	"jv/team-tone-tuner/utils"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type (
	JWTConfig struct {
		Skipper    Skipper
		SigningKey interface{}
	}
	Skipper      func(c echo.Context) bool
	jwtExtractor func(echo.Context) (string, error)
)

const (
	JWT_COOKIE_NAME = "jwt"
)

var (
	ErrJWTMissing = echo.NewHTTPError(http.StatusUnauthorized, "missing or malformed jwt")
	ErrJWTInvalid = echo.NewHTTPError(http.StatusForbidden, "invalid or expired jwt")
)

func JWT(key []byte) echo.MiddlewareFunc {
	c := JWTConfig{}
	c.SigningKey = key
	return JWTWithConfig(c)
}

func JWTWithConfig(config JWTConfig) echo.MiddlewareFunc {
	extractor := jwtFromCookie(JWT_COOKIE_NAME)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth, err := extractor(c)

			if err != nil {
				if config.Skipper != nil {
					if config.Skipper(c) {
						return next(c)
					}
				}
				return c.JSON(http.StatusUnauthorized, utils.NewError(err))
			}

			token, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
				}
				return config.SigningKey, nil
			})

			if err != nil {
				return c.JSON(http.StatusForbidden, utils.NewError(ErrJWTInvalid))
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				userID := claims["id"].(string)
				roleString := claims["role"].(string)
				// We need to convert the string to the UserRole type
				role, ok := utils.UserRoleLookup[roleString]

				if !ok {
					return c.JSON(http.StatusForbidden, utils.NewError(ErrJWTInvalid))
				}

				uuid, err := uuid.Parse(userID)

				if err != nil {
					return c.JSON(http.StatusForbidden, utils.NewError(ErrJWTInvalid))
				}

				c.Set("userId", uuid)
				c.Set("role", role)

				return next(c)
			}

			return c.JSON(http.StatusForbidden, utils.NewError(ErrJWTInvalid))
		}
	}
}

func jwtFromCookie(key string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		cookie, err := c.Cookie(key)

		if err != nil {
			return "", ErrJWTMissing
		}

		return cookie.Value, nil
	}
}
