package vo

import (
	"errors"
	"regexp"
)

type Mobile struct {
	Value string `binding:"number,mobile"`
}

func (o Mobile) Validate() error {
	if !regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(o.Value) {
		return errors.New("手机号码格式不正确")
	}
	return nil
}
