package user

import "github.com/tidwall/gjson"

type Age struct {
	Value int64 `binding:"min=1,max=150"`
	Set   bool
}

func NewAge(v int64) Age {
	return Age{Value: v, Set: true}
}

func (age Age) IsSet() bool {
	return age.Set
}

func (age Age) IsPresent(f func(v int64)) {
	if age.Set {
		f(age.Value)
	}
}

func (age *Age) UnmarshalJSON(data []byte) error {
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		age.Value = result.Int()
		age.Set = true
	}
	return nil
}
