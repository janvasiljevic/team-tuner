package middleware

import (
	"jv/team-tone-tuner/utils"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	ErrUserRole = echo.NewHTTPError(http.StatusForbidden, "invalid user role")
)

func CheckUserRole(role utils.UserRole) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			currentRole := c.Get("role").(utils.UserRole)

			if currentRole != role {
				userId := c.Get("userId").(uuid.UUID)
				log.Warn().Str("userId", userId.String()).Msg("Invalid user role")
				return c.JSON(http.StatusForbidden, utils.NewError(ErrUserRole))
			}

			return next(c)
		}
	}
}
