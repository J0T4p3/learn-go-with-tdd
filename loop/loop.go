package loop

import "strings"

const (
	upper = "upper"
	lower = "lower"
)

func Loop(item string, times int, wcase string) string {
	var looped strings.Builder

	for range times {
		looped.WriteString(item)
	}
	finalString := looped.String()

	switch wcase {
	case upper:
		finalString = strings.ToUpper(finalString)
	case lower:
		finalString = strings.ToLower(finalString)
	}

	return finalString
}
