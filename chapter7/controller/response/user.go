package response

import "time"

type Users []User

type User struct {
	ID         *int64     `json:",omitempty"`
	HashID     *string    `json:",omitempty"`
	Mobile     *string    `json:",omitempty"`
	Age        *int64     `json:",omitempty"`
	Level      *int64     `json:",omitempty"`
	LevelDesc  *string    `json:",omitempty"`
	Nickname   *string    `json:",omitempty"`
	Avatar     *string    `json:",omitempty"`
	Bio        *string    `json:",omitempty"`
	Amount     *int64     `json:",omitempty"`
	CreateTime *time.Time `json:",omitempty"`
}
