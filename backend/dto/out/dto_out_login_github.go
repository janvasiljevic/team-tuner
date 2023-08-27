package out

type GithubLoginOut struct {
	RedirectUrl string `json:"redirectUrl" validate:"required"`
}
