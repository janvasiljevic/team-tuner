package schema

import (
	"jv/team-tone-tuner/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BfiAnswer holds the schema definition for the BfiAnswer entity.
type BfiAnswer struct {
	ent.Schema
}

func (BfiAnswer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the BfiAnswer.
func (BfiAnswer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value").Optional().Nillable().Min(1).Max(5),
	}
}

// Edges of the BfiAnswer.
func (BfiAnswer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("bfi_question", BfiQuestion.Type).Ref("bfi_answers").Unique(),
		edge.From("student", User.Type).Ref("bfi_answers").Unique(),
	}
}
