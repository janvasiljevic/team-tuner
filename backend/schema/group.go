package schema

import (
	"jv/team-tone-tuner/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Student entity.
type Group struct {
	ent.Schema
}

func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		// The name of the group
		field.String("name").NotEmpty(),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		// Students who belong to the group
		edge.To("students", User.Type),

		// The course the group belongs to
		edge.From("course", Course.Type).Ref("groups").Unique().Required(),

		// The group run the group belongs to
		edge.From("group_run", GroupRun.Type).Ref("groups").Unique().Required(),
	}
}
