package user

import "github.com/tidwall/gjson"

type Age struct {
	Value int64 `binding:"min=1,max=150"`
	Set   bool
	Mask  bool
}

func NewAge(v int64) Age {
	return Age{Value: v, Set: true}
}

func (age Age) IsMask() bool {
	return age.Mask
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
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		age.Value = results[0].Int()
		age.Set = true
	}
	if results[1].Exists() {
		age.Mask = results[1].Bool()
	}
	return nil
}
