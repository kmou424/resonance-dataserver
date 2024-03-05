package strkit

import "strings"

func Concat(ss ...string) string {
	var builder strings.Builder
	for _, s := range ss {
		builder.WriteString(s)
	}
	return builder.String()
}
