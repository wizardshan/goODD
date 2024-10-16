package user

import (
	"errors"
	"github.com/tidwall/gjson"
	"goODD/chapter5/pkg/validator"
)

type NicknameV2 struct {
	Value   string
	Keyword *string
}

func (o NicknameV2) validate(v string) error {
	return validator.Var(v, "min=2,max=10")
}

func (o NicknameV2) Validate() error {
	return o.validate(o.Value)
}

func (o NicknameV2) ValidateOmit() error {
	if o.Value != "" {
		return o.validate(o.Value)
	}
	return nil
}

func (o NicknameV2) FuzzyValidate() error {
	if o.Keyword == nil || *o.Keyword == "" {
		return errors.New("模糊关键词不能为空")
	}
	return o.validate(*o.Keyword)
}

func (o NicknameV2) FuzzyValidateOmit() error {
	if o.Keyword == nil {
		return nil
	}
	if *o.Keyword == "" {
		return errors.New("模糊关键词不能为空")
	}
	return o.validate(*o.Keyword)
}

func (o *NicknameV2) UnmarshalJSON(data []byte) error {
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
		keyword := results[1].String()
		o.Keyword = &keyword
	}
	return nil
}
