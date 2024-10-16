package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter5/pkg/validator"
)

type Nickname struct {
	Value   string
	Keyword string
}

func (o Nickname) validate(v string) error {
	return validator.Var(v, "min=2,max=10")
}

func (o Nickname) Validate() error {
	return o.validate(o.Value)
}

func (o Nickname) ValidateOmit() error {
	if o.Value != "" {
		return o.validate(o.Value)
	}
	return nil
}

func (o Nickname) FuzzyValidate() error {
	if o.Keyword == "" {
		return errors.New("模糊关键词不能为空")
	}
	return o.validate(o.Keyword)
}

func (o Nickname) FuzzyValidateOmit() error {
	if o.Keyword == "" {
		return errors.New("模糊关键词不能为空")
	}
	return o.validate(o.Keyword)
}

func (o *Nickname) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).String()
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Keyword")
	if results[0].Exists() {
		o.Value = results[0].String()
		return nil
	}
	if results[1].Exists() {
		o.Keyword = results[1].String()
	}
	return nil
}
