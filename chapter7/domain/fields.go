package domain

import (
	"github.com/samber/lo"
	"strings"
)

type Fields map[string]string

func (fields Fields) Names(key string) ([]string, bool) {
	namesBySep, ok := fields[lo.PascalCase(key)]
	if !ok {
		return nil, ok
	}
	return strings.Split(namesBySep, ","), true
}

func (fields Fields) SnakeCaseNames(key string) ([]string, bool) {
	names, ok := fields.Names(key)
	if !ok {
		return nil, ok
	}
	return lo.Map(names, func(item string, index int) string {
		return lo.SnakeCase(item)
	}), true
}
