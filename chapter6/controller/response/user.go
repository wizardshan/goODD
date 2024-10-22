package response

import (
	"time"
)

type Users []User

type User struct {
	ID         int64
	HashID     string
	Mobile     string
	Age        int64
	Level      int64
	LevelDesc  string
	Nickname   string
	Avatar     string
	Bio        string
	Amount     float64
	CreateTime time.Time
}
