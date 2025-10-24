package iteration

import "strings"

func Repeat(input string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(input)
	}
	return repeated.String()
}
