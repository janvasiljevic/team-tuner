package api

import (
	"jv/team-tone-tuner/dto/out"
	"jv/team-tone-tuner/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//	@Summary		Get who am I
//	@Description	Get who am I: Information about the current user
//	@Tags			auth
//	@Success		200	{object}	out.WhoAmIOut
//	@Router			/auth/whoami [get]
func (api Api) GetWhoAmI(c echo.Context) error {

	userId := c.Get("userId").(uuid.UUID)

	user, err := api.userStore.GetUserInfo(userId, c.Request().Context())

	if err != nil {
		return err
	}

	if user == nil {
		return c.String(404, "User not found")
	}

	return c.JSON(200, out.NewWhoAmIOut(user))
}

//	@Summary		Logout
//	@Description	Logout the current user: Clears the JWT cookie
//	@Tags			auth
//	@Success		200
//	@Router			/auth/logout [post]
func (api Api) Logout(c echo.Context) error {

	c.SetCookie(utils.ClearJWtCookie())

	return c.NoContent(200)
}
