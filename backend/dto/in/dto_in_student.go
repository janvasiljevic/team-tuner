package in

import (
	"jv/team-tone-tuner/dto"

	"github.com/google/uuid"
)

type StudentSortField string

const (
	StudentSortFieldOpenness      StudentSortField = "openness"
	StudentSortFieldConscientious StudentSortField = "conscientious"
	StudentSortFieldExtraversion  StudentSortField = "extraversion"
	StudentSortFieldAgreeableness StudentSortField = "agreeableness"
	StudentSortFieldNeuroticism   StudentSortField = "neuroticism"
)

type GetStudentsQuery struct {
	CourseId            *uuid.UUID        `query:"courseId" validate:"omitempty,uuid4"`
	CompletedQuestioner *bool             `query:"completedQuestioner"`
	SortField           *StudentSortField `query:"sortField"`
	SortOrder           *dto.SortOrder    `query:"sortOrder"`
	PageSize            *int              `query:"pageSize"`
	Page                *int              `query:"page"`
}
