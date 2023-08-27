package out

import (
	"jv/team-tone-tuner/model"
	"jv/team-tone-tuner/model/user"
)

type OutUserRole string

const (
	RoleStudent OutUserRole = "student"
	RoleTeacher OutUserRole = "admin"
)

type WhoAmIOut struct {
	Id             string      `json:"id" validate:"required"`
	GithubUsername string      `json:"github_username" validate:"required"`
	OutUserRole    OutUserRole `json:"role" validate:"required"`
	FinishedBfi    bool        `json:"finished_bfi" validate:"required"`
}

func NewWhoAmIOut(modelIn *model.User) WhoAmIOut {
	out := WhoAmIOut{
		Id:             modelIn.ID.String(),
		GithubUsername: modelIn.GithubUsername,
		FinishedBfi:    modelIn.Edges.BfiReport != nil,
	}

	if modelIn.Role == user.RoleAdmin {
		out.OutUserRole = RoleTeacher
	}

	if modelIn.Role == user.RoleStudent {
		out.OutUserRole = RoleStudent
	}

	return out
}
