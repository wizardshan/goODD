package vo

import (
	"errors"
	"regexp"
)

type Mobile struct {
	StringValue
}

func (o *Mobile) Mask() string {
	return o.Value[0:4] + "****" + o.Value[8:11]
}

func (o *Mobile) GenNickname() string {
	return o.Mask()
}

func (o *Mobile) ValidateOmit() error {
	if o.Set {
		return o.Validate()
	}
	return nil
}

func (o *Mobile) Validate() error {
	if !regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(o.Value) {
		return errors.New("手机号码格式不正确")
	}
	return nil
}
