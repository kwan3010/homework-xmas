package schema

import "github.com/facebookincubator/ent"

// Watcher holds the schema definition for the Watcher entity.
type Watcher struct {
	ent.Schema
}

// Fields of the Watcher.
func (Watcher) Fields() []ent.Field {
	return nil
}

// Edges of the Watcher.
func (Watcher) Edges() []ent.Edge {
	return nil
}
