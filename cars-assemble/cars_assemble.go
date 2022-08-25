package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * (successRate / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	minutesPerHour := 60
	return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / minutesPerHour
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	carsPerGroup := 10
	carsGroupCost := 95000
	individualCarCost := 10000

	carGroupsCount := carsCount / carsPerGroup
	individualCarsCount := carsCount % carsPerGroup

	return uint(carGroupsCount*carsGroupCost + individualCarsCount*individualCarCost)
}
