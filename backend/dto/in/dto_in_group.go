package in

import (
	"github.com/google/uuid"
)

type GetGroupsQuery struct {
	GroupRun *uuid.UUID `query:"groupRun" validate:"omitempty,uuid4"`
}

type PostFormGroupsIn struct {
	// Array length must be at least 4
	// Each element must be greater than 0 and less than 10
	GroupSizes []int `body:"groupSizes" validate:"required,min=4,dive,gt=0,lt=10"`

	// General settings for SA
	Iterations  int     `body:"iterations" validate:"required,gt=0"`
	Temperature float64 `body:"temperature" validate:"required,gt=0"`

	// Weights for SA
	WeightSatisfaction      float64 `body:"weightSatisfaction" validate:"required,gt=0,lt=3"`
	WeightNeuroticism       float64 `body:"weightNeuroticism" validate:"required,gt=0,lt=3"`
	WeightExtraversion      float64 `body:"weightExtraversion" validate:"required,gt=0,lt=3"`
	WeightConscientiousness float64 `body:"weightConscientiousness" validate:"required,gt=0,lt=3"`

	// To which course the groups belong
	CourseId *uuid.UUID `param:"courseId" validate:"required,uuid4"`
}
