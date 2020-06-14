package internal

import (
	"strings"
	"unicode"
)

func CamelCase(s string) string {
	o := ""

	for _, n := range strings.Split(strings.ToLower(s), "_") {
		o += strings.Title(n)
	}

	return o
}

func SnakeCase(s string) string {
	o := []rune{}

	for i, r := range []rune(s) {
		if unicode.IsUpper(r) {
			if i > 0 {
				o = append(o, '_')
			}

			o = append(o, unicode.ToLower(r))
		} else {
			o = append(o, r)
		}
	}

	return string(o)
}
