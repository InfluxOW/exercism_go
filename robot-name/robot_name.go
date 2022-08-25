package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var alphabet = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var alphabetLength = len(alphabet)

var codeLettersPartCapacity = alphabetLength * alphabetLength
var codeNumbersPartCapacity = 10 * 10 * 10

var availableCodes []int

func init() {
	rand.Seed(time.Now().UnixNano())

	availableCodes = rand.Perm(codeLettersPartCapacity * codeNumbersPartCapacity)
}

type NamedCode struct {
	code int
}

func (nc NamedCode) String() string {
	lp := nc.code / codeNumbersPartCapacity
	np := nc.code % codeNumbersPartCapacity

	l1 := lp / alphabetLength
	l2 := lp % alphabetLength

	return fmt.Sprintf("%s%s%.3d", alphabet[l1], alphabet[l2], np)
}

type Robot struct {
	name  NamedCode
	named bool
}

func (r *Robot) Name() (string, error) {
	if r.named {
		return r.name.String(), nil
	}

	err := r.Reset()

	if err != nil {
		return "", err
	}

	return r.name.String(), nil
}

func (r *Robot) Reset() error {
	if len(availableCodes) == 0 {
		return errors.New("name: no available names left")
	}

	code := availableCodes[0]

	availableCodes = availableCodes[1:]
	r.name = NamedCode{code: code}
	r.named = true

	return nil
}
