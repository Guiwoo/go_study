package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"semina_entgo/custom"
)

// Tester holds the schema definition for the Tester entity.
type Tester struct {
	ent.Schema
}

// Fields of the Tester.
func (Tester) Fields() []ent.Field {
	return []ent.Field{
		field.String("PascalCase"),
		field.String("let_me_check"),
		field.Enum("size").
			Values("big", "small"),
		field.Enum("shape").
			GoType(custom.Shape("")),
		field.Enum("level").
			GoType(custom.Level(0)),
	}
}

// Edges of the Tester.
func (Tester) Edges() []ent.Edge {
	return nil
}
