package schema

import (
	"errors"
	"jv/team-tone-tuner/schema/mixins"
	"regexp"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

func (Course) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.BaseMixin{},
	}
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),          // Osnove Podatkovnih baz
		field.String("code").Unique(), // OPB
		field.String("colour").Default("#000000").Validate(func(s string) error {
			hexColorPattern := "^#([A-Fa-f0-9]{6}|[A-Fa-f0-9]{3})$"
			hexColorRegex := regexp.MustCompile(hexColorPattern)

			if hexColorRegex.MatchString(s) {
				return nil
			}

			return errors.New("colour must be a valid hex colour")
		}),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		// One course can have many students
		edge.To("students", User.Type),
		edge.To("groups", Group.Type),

		edge.To("group_runs", GroupRun.Type),
	}
}
