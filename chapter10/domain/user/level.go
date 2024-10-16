package user

import (
	"goODD/chapter10/domain/vo"
	"goODD/chapter10/pkg/validator"
)

const (
	LevelNormal   = 0
	LevelSilver   = 10
	LevelGold     = 20
	LevelPlatinum = 30
)

type Level struct {
	vo.Int64Multi
	Desc string
}

func (o Level) GetDesc() string {
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

func (o Level) validate(v int64) error {
	return validator.Var(v, "oneof=0 10 20 30")
}

func (o Level) Validate() error {
	return o.validate(o.Value)
}

func (o Level) ValidateOmit() error {
	if o.Set {
		return o.validate(o.Value)
	}
	return nil
}

func (o Level) MultiValidate() error {
	return o.Int64Multi.MultiValidate(o.validate)
}

func (o Level) MultiValidateOmit() error {
	return o.Int64Multi.MultiValidateOmit(o.validate)
}
