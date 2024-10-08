package domain

import (
	"regexp"
)

type Mobile struct {
	Value string `binding:"number,mobile"`
}

func (o Mobile) IsValid() bool {
	return regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(o.Value)
}
