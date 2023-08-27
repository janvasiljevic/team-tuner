package schema

import (
	"jv/team-tone-tuner/schema/mixins"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupRun holds the schema definition for the Student entity.
type GroupRun struct {
	ent.Schema
}

func (GroupRun) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the GroupRun.
func (GroupRun) Fields() []ent.Field {
	return []ent.Field{
		field.Ints("group_size"),
		field.Floats("best_score_history"),
		field.Floats("candidate_score_history"),
	}

}

// Edges of the GroupRun.
func (GroupRun) Edges() []ent.Edge {
	return []ent.Edge{
		// Who created the group
		edge.To("created_by", User.Type).Unique().Required(),

		// Which groups were created in this run
		edge.To("groups", Group.Type),

		// The course the group belongs to
		edge.From("course", Course.Type).Ref("group_runs").Unique().Required(),
	}
}
