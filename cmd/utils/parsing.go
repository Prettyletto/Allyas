package utils

import "strings"

func Attach(builder *strings.Builder, parts ...string) {
	for _, v := range parts {
		builder.WriteString(v)
	}
}
