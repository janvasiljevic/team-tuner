package schema

import (
	"jv/team-tone-tuner/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("github_username").Unique(),
		field.String("univeresity_id").Unique().Nillable().Optional(),
		field.Enum("role").Values("student", "admin").Default("student"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// student: can have many courses
		edge.From("courses", Course.Type).Ref("students"),

		// student: specifc report
		edge.From("bfi_report", BfiReport.Type).Ref("student").Unique(),

		// student: specifics answers
		edge.To("bfi_answers", BfiAnswer.Type),

		// student: can be in many groups
		edge.From("groups", Group.Type).Ref("students"),
	}
}
