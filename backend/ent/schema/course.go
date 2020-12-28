package schema

import "github.com/facebookincubator/ent"

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
//อ2
func (Course) Fields() []ent.Field {
	return nil
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return nil
}
