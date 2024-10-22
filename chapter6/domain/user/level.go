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

type Levels []Level

func (o Levels) IsPresent(f func(v []int64)) {
	if len(o) != 0 {
		f(o.Values())
	}
}

func (o Levels) Values() []int64 {
	values := make([]int64, len(o))
	for i, level := range o {
		values[i] = level.Value
	}
	return values
}

type Level struct {
	Value int64 `binding:"oneof=0 10 20 30"`
	Set   bool
}

func (o *Level) Desc() string {
	mapping := map[int64]string{
		LevelNormal:   "普通",
		LevelSilver:   "白银",
		LevelGold:     "黄金",
		LevelPlatinum: "铂金",
	}
	desc, ok := mapping[o.Value]
	if ok {
		return desc
	}
	return "未知"
}

func (o *Level) IsPresent(f func(v int64)) {
	if o.Set {
		f(o.Value)
	}
}

func (o *Level) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

func (o *Level) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.SetTo(gjson.ParseBytes(data).Int())
		return nil
	}

	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		o.SetTo(result.Int())
	}
	return nil
}
