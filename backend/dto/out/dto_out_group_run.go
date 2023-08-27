package out

import (
	"jv/team-tone-tuner/model"
	"time"
)

type GroupRunOut struct {
	ID        string    `json:"id" validate:"required"`
	CreatedAt time.Time `json:"createdAt" validate:"required"`
	UpdatedAt time.Time `json:"updatedAt" validate:"required"`
}

func ConvertModelToGroupRunOut(model *model.GroupRun) GroupRunOut {
	return GroupRunOut{
		ID:        model.ID.String(),
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}
