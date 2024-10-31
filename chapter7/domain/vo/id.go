package vo

import (
	"goODD/chapter5/pkg/validate"
)

type ID struct {
	Int64Value
}

func (o *ID) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *ID) Validate() error {
	return validate.Var(o.Value, "min=1")
}
