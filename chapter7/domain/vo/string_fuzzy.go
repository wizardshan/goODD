package vo

import (
	"errors"
	"github.com/tidwall/gjson"
)

type StringFuzzy struct {
	StringValue
	Keyword StringValue
}

func (o StringFuzzy) FuzzyIsPresent(f func(v string)) {
	if o.Keyword.Set {
		f(o.Keyword.Value)
	}
}

func (o StringFuzzy) FuzzyEmpty() bool {
	return o.Keyword.Value == ""
}

func (o StringFuzzy) FuzzyValidate(validate func(v string) error) error {
	if o.FuzzyEmpty() {
		return errors.New("模糊关键词不能为空")
	}
	return validate(o.Keyword.Value)
}

func (o StringFuzzy) FuzzyValidateOmit(validate func(v string) error) error {
	if !o.Keyword.Set {
		return nil
	}
	if o.FuzzyEmpty() {
		return errors.New("模糊关键词不能为空")
	}
	return validate(o.Keyword.Value)
}

func (o *StringFuzzy) UnmarshalJSON(data []byte) error {
	if data[0] != '{' {
		o.Value = gjson.ParseBytes(data).String()
		o.Set = true
		return nil
	}

	results := gjson.GetManyBytes(data, "Value", "Keyword")
	if results[0].Exists() {
		o.SetTo(results[0].String())
		return nil
	}
	if results[1].Exists() {
		o.Keyword.SetTo(results[1].String())
	}
	return nil
}
