package iteration

import "strings"

const repeatCount = 5

func Repeat(character string) string {
	return RepeatTimes(character, repeatCount)
}

func RepeatTimes(character string, repeatTimes int) string {
	var repeated strings.Builder
	for range repeatTimes {
		repeated.WriteString(character)
	}
	return repeated.String()
}
