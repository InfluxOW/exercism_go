package clock

import "fmt"

const (
	mPerH = 60
	hPerD = 24
	mPerD = hPerD * mPerH
)

type Clock struct {
	h uint
	m uint
}

func New(h int, m int) Clock {
	m = h*mPerH + m
	if m < 0 {
		m = mPerD + m%mPerD
	}

	return Clock{
		h: uint((m / mPerH) % hPerD),
		m: uint(m % mPerH),
	}
}

func (c Clock) Add(m int) Clock {
	return New(int(c.h), int(c.m)+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(int(c.h), int(c.m)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.h, c.m)
}
