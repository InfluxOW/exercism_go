package thefarm

import (
	"errors"
	"fmt"
)

type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

var (
	negativeFodderErr = errors.New("negative fodder")
	divisionByZeroErr = errors.New("division by zero")
)

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	am, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction {
		am *= 2
		err = nil
	}

	switch {
	case err != nil:
	case am < 0:
		err = negativeFodderErr
	case cows == 0:
		err = divisionByZeroErr
	case cows < 0:
		err = &SillyNephewError{cows: cows}
	}

	if err != nil {
		return 0, err
	}

	return am / float64(cows), nil
}
