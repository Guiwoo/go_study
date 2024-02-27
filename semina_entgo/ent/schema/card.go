package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"regexp"
	"time"
)

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			Match(regexp.MustCompile("([0-9]{4})-([0-9]{4})-([0-9]{4})-([0-9]{4})")),
		field.Time("expired_at").Default(time.Now()),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("owner", User.Type).Unique(),
	}
}
