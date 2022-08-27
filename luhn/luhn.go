package luhn

import (
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	ln := len(id)

	if ln < 2 {
		return false
	}

	isEven := ln%2 == 0
	isSecond := func(i int) bool {
		if isEven {
			return i%2 == 0
		}
		return (i+1)%2 == 0
	}
	isNotDigit := func(n int32) bool {
		return n > '9' || n < '0'
	}

	sum := 0
	for i, n := range id {
		if isNotDigit(n) {
			return false
		}

		n := int(n - '0')

		if isSecond(i) {
			n *= 2

			if n > 9 {
				n -= 9
			}
		}

		sum += n
	}

	return sum%10 == 0
}
