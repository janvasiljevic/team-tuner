package schema

import (
	"jv/team-tone-tuner/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BfiQuestion holds the schema definition for the BfiQuestion entity.
type BfiQuestion struct {
	ent.Schema
}

func (BfiQuestion) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the BfiQuestion.
func (BfiQuestion) Fields() []ent.Field {
	return []ent.Field{
		field.String("questiono").NotEmpty(),
		field.String("facet").NotEmpty(),
		field.Enum("dimension").Values("extraversion", "agreeableness", "conscientiousness", "neuroticism", "openness"),
		field.Enum("influence").Values("positive", "negative"),
		field.Float("alpha").Positive().Min(0).Max(1),
	}
}

// Edges of the BfiQuestion.
func (BfiQuestion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("bfi_answers", BfiAnswer.Type),
	}
}
