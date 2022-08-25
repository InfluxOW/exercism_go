package birdwatcher

const (
	daysPerWeek = 7
)

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	r := 0
	for _, b := range birdsPerDay {
		r += b
	}
	return r
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birdsPerGivenWeek []int
	if week == 1 {
		birdsPerGivenWeek = birdsPerDay[:week*daysPerWeek]
	} else {
		birdsPerGivenWeek = birdsPerDay[(week-1)*daysPerWeek : week*daysPerWeek]
	}

	return TotalBirdCount(birdsPerGivenWeek)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, b := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = b + 1
		} else {
			birdsPerDay[i] = b
		}
	}
	return birdsPerDay
}
