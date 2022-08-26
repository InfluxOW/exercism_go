package expenses

import (
	"fmt"
)

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	var rs []Record
	for _, r := range in {
		if predicate(r) {
			rs = append(rs, r)
		}
	}
	return rs
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(p DaysPeriod) func(Record) bool {
	return func(r Record) bool {
		return r.Day >= p.From && r.Day <= p.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(c string) func(Record) bool {
	return func(r Record) bool {
		return r.Category == c
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, p DaysPeriod) float64 {
	return sumExpenses(Filter(in, ByDaysPeriod(p)))
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, p DaysPeriod, c string) (float64, error) {
	fIn := Filter(in, ByCategory(c))
	if len(fIn) == 0 && len(fIn) < len(in) {
		return 0.0, fmt.Errorf("unknown category %s", c)
	}
	fIn = Filter(fIn, ByDaysPeriod(p))

	return sumExpenses(fIn), nil
}

func sumExpenses(in []Record) float64 {
	e := 0.0
	for _, r := range in {
		e += r.Amount
	}
	return e
}
