package api

import (
	"fmt"
	"jv/team-tone-tuner/config"
	"jv/team-tone-tuner/dto/in"
	"jv/team-tone-tuner/dto/out"
	"jv/team-tone-tuner/model/user"
	"jv/team-tone-tuner/service"
	"jv/team-tone-tuner/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

//	@Summary		Github login
//	@Description	Github login
//	@Tags			login
//	@Param			body	body		in.GithubLoginIn	true	"Github login body"
//	@Success		200		{object}	out.GithubLoginOut
//	@Router			/login/github [post]
func (api Api) PostGithubLogin(c echo.Context) error {
	var dto in.GithubLoginIn

	if err := c.Bind(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.BadRequest())
	}

	// validate
	if err := c.Validate(&dto); err != nil {
		return c.JSON(http.StatusBadRequest, out.NewValidatorError(err))
	}

	redirectUrl := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s&state=%s",
		config.LoadedConfig.Github.ClientId,
		config.LoadedConfig.Github.RedirectUrl,
		dto.NavigateTo,
	)

	return c.JSON(http.StatusOK, out.GithubLoginOut{RedirectUrl: redirectUrl})
}

// Returned from github login callback
// Not exposed to the Swagger API, because it's not meant to be called by the frontend
func (api Api) GetGithubLoginCallback(c echo.Context) error {
	ctx := c.Request().Context()

	code := c.QueryParam("code")
	redirectUrl := c.QueryParam("state")

	githubAccessToken, err := service.GetGithubAccessToken(code)

	if err != nil {
		log.Error().Err(err).Msg("Error getting github access token")

		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	githubUsername, err := service.GetGithubUsername(githubAccessToken)

	if err != nil {
		log.Error().Err(err).Msg("Error getting github username")

		return c.JSON(http.StatusInternalServerError, out.NewInternal())
	}

	foundUser := api.db.User.Query().Where(user.GithubUsernameEQ(githubUsername)).FirstX(ctx)

	if foundUser != nil {
		var role utils.UserRole

		if foundUser.Role == user.RoleAdmin {
			role = utils.UserRoleAdmin
		} else {
			role = utils.UserRoleStudent
		}

		c.SetCookie(utils.GenerateJWTCookie(foundUser.ID, role))

		return c.Redirect(http.StatusFound, redirectUrl)
	}

	newStudent := api.userStore.CreateANewStudentWithAnswers(githubUsername, ctx)

	c.SetCookie(utils.GenerateJWTCookie(newStudent.ID, utils.UserRoleStudent))

	// If we get here, the user is not in the database, dont set a cookie
	return c.Redirect(http.StatusFound, redirectUrl)
}
