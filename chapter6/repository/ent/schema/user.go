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
		field.String("mobile").Default(""),
		field.Int64("age").Default(0),
		field.Int64("level").Default(0),
		field.String("nickname").Default(""),
		field.Time("create_time").Default(time.Now),
	}
}
