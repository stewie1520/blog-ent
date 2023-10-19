package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().StructTag(`json:"id"`),
		field.String("email").NotEmpty().StructTag(`json:"email"`).Unique(),
		field.String("password").NotEmpty().StructTag(`json:"password"`),
		field.Time("created_at").Immutable().Default("now()").StructTag(`json:"created_at"`),
		field.Time("updated_at").Default("now()").UpdateDefault("now()").StructTag(`json:"updated_at"`),
		field.Time("deleted_at").Optional().Nillable().StructTag(`json:"deleted_at"`),
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("account").Unique().StructTag(`json:"user"`),
	}
}
