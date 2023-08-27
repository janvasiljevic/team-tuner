package in

type GithubLoginIn struct {
	NavigateTo string `json:"navigateTo" validate:"required,url"`
}
