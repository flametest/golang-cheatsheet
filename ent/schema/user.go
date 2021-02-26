package schema

import (
	"encoding/json"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func RawJSON(name string) ent.Field {
	return field.Bytes(name).SchemaType(map[string]string{
		dialect.MySQL:    "JSON",
		dialect.Postgres: "JSON",
	}).GoType(json.RawMessage{})
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int64("money").Default(0),
		RawJSON("meta"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "User"},
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
