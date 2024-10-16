package response

import (
	"time"
)

type Users []User

type User struct {
	ID         *int64     `json:",omitempty"`
	Mobile     *string    `json:",omitempty"`
	Age        *int64     `json:",omitempty"`
	Level      *int64     `json:",omitempty"`
	Nickname   *string    `json:",omitempty"`
	CreateTime *time.Time `json:",omitempty"`
}
