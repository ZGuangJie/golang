package testgo

import (
	"strings"
)

func Split(s, sep string) (result []string) {
	result = make([]string, 0, strings.Count(s, sep)+1)
	index := strings.Index(s, sep)
	for index > -1 {
		result = append(result, s[:index])
		s = s[index+len(sep):]
		index = strings.Index(s, sep)
	}
	result = append(result, s)

	return
}
