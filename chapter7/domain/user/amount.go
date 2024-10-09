package user

import "github.com/tidwall/gjson"

type Amount struct {
	Value int64
	Set   bool
}

func NewAmount(v int64) Amount {
	return Amount{Value: v, Set: true}
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
	result := gjson.GetBytes(data, "Value")
	if result.Exists() {
		amount.Value = result.Int()
		amount.Set = true
	}
	return nil
}
