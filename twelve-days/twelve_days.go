package twelve

import (
	"fmt"
	"strings"
)

var intToWord = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var intToGift = map[int]string{
	1:  "a Partridge in a Pear Tree.",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

func Verse(i int) string {
	var gifts string
	for n := i; n > 0; n-- {
		if len(gifts) > 0 {
			if n == 1 {
				gifts += ", and "
			}

			if n > 1 {
				gifts += ", "
			}
		}

		gifts += intToGift[n]
	}

	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s", intToWord[i], gifts)
}

func Song() string {
	var vs []string
	for i := 1; i <= len(intToGift); i++ {
		vs = append(vs, Verse(i))
	}

	return strings.Join(vs, "\n")
}
