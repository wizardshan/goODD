package user

import (
	"github.com/tidwall/gjson"
)

const (
	LevelNormal   = 0
	LevelSilver   = 10
	LevelGold     = 20
	LevelPlatinum = 30
)

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
	Desc  string
	Set   bool
}

func NewLevel(v int64) Level {
	level := Level{Value: v, Set: true}
	level.Desc = level.GetDesc()
	return level
}

func (level Level) GetDesc() string {
	desc, ok := level.DescMapping()[level.Value]
	if ok {
		return desc
	}
	return "未知"
}

func (level Level) DescMapping() map[int64]string {
	return map[int64]string{
		LevelNormal:   "普通",
		LevelSilver:   "白银",
		LevelGold:     "黄金",
		LevelPlatinum: "铂金",
	}
}

func (level Level) IsSet() bool {
	return level.Set
}

func (level Level) IsPresent(f func(v int64, desc string)) {
	if level.Set {
		f(level.Value, level.Desc)
	}
}

func (level *Level) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		level.Value = result.Int()
		level.Set = true
	}
	return nil
}
