package domain

import (
	"regexp"
)

type Mobile struct {
	Value string `binding:"number,mobile"`
	Set   bool
}

func (mobile Mobile) IsValid() bool {
	return regexp.MustCompile(`^(1[3-9][0-9]\d{8})$`).MatchString(mobile.Value)
}
