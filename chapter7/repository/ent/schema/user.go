package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("hash_id").Default(""),
		field.String("mobile").Default(""),
		field.String("password").Default(""),
		field.Int64("age").Default(0),
		field.Int64("level").Default(0),
		field.String("nickname").Default(""),
		field.String("avatar").Default(""),
		field.String("bio").Default(""),
		field.Int64("amount").Default(0),
		field.Int64("status").Default(0),
		field.Time("create_time").Default(time.Now),
		field.Time("update_time").Default(time.Now).UpdateDefault(time.Now),
	}
}
