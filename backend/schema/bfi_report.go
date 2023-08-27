package schema

import (
	"jv/team-tone-tuner/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type BfiReportItem struct {
	PointsSum        int     `json:"pointsSum" validate:"required"`
	PointsMax        int     `json:"pointsMax" validate:"required"`
	PointsMin        int     `json:"pointsMin" validate:"required"`
	PointsNormalized float64 `json:"pointsNormalized" validate:"required"`
}

// BfiReport holds the schema definition for the BfiReport entity.
type BfiReport struct {
	ent.Schema
}

func (BfiReport) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the BfiReport.
func (BfiReport) Fields() []ent.Field {
	return []ent.Field{
		field.JSON("conscientiousness", BfiReportItem{}),
		field.JSON("extraversion", BfiReportItem{}),
		field.JSON("agreeableness", BfiReportItem{}),
		field.JSON("neuroticism", BfiReportItem{}),
		field.JSON("openness", BfiReportItem{}),
	}
}

// Edges of the BfiReport.
func (BfiReport) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("student", User.Type).Unique(),
	}
}
