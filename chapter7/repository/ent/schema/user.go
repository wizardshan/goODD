package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("hash_id").Default(""),
		field.String("mobile").Default(""),
		field.String("password").Sensitive(),
		field.Int64("age").Default(0),
		field.Int64("level").Default(0),
		field.String("nickname").Default(""),
		field.String("avatar").Default(""),
		field.String("bio").Default(""),
		field.Int64("amount").Default(0),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (User) Edges() []ent.Edge {
	return nil
}
