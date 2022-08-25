package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	f := make(Frequency)
	for _, w := range regexp.MustCompile(`\w+('\w+)?`).FindAllString(strings.ToLower(phrase), -1) {
		f[w]++
	}
	return f
}
