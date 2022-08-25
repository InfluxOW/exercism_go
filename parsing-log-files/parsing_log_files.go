package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	matched, err := regexp.MatchString(`^\[(TRC|DBG|INF|WRN|ERR|FTL)]`, text)

	return matched && err == nil
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i) password`)

	c := 0
	for _, l := range lines {
		if re.MatchString(l) {
			c++
		}
	}

	return c
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d*`)

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w*)`)

	var r []string
	for _, l := range lines {
		f := re.FindStringSubmatch(l)
		if len(f) > 1 {
			u := f[1]
			l = fmt.Sprintf("[USR] %s %s", u, l)
		}

		r = append(r, l)
	}

	return r
}
