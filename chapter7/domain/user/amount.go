package user

import "github.com/tidwall/gjson"

type Amount struct {
	Value int64
	Set   bool
	Mask  bool
}

func NewAmount(v int64) Amount {
	return Amount{Value: v, Set: true}
}

func (amount Amount) IsMask() bool {
	return amount.Mask
}

func (amount Amount) IsSet() bool {
	return amount.Set
}

func (amount Amount) IsPresent(f func(v int64)) {
	if amount.Set {
		f(amount.Value)
	}
}

func (amount *Amount) UnmarshalJSON(data []byte) error {
	results := gjson.GetManyBytes(data, "Value", "Mask")
	if results[0].Exists() {
		amount.Value = results[0].Int()
		amount.Set = true
	}
	if results[1].Exists() {
		amount.Mask = results[1].Bool()
	}
	return nil
}
