package raindrops

import "strconv"

var factorsToSounds map[int]string = map[int]string{
	3: "Pling",
	5: "Plang",
	7: "Plong",
}

var orderedFactors []int = []int{3, 5, 7}

func Convert(number int) string {
	var r string

	for _, f := range orderedFactors {
		if number%f == 0 {
			r += factorsToSounds[f]
		}
	}

	if r == "" {
		r = strconv.Itoa(number)
	}

	return r
}
