package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	rate := float64(productionRate)
	return (rate * successRate) / 100
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	rate := float64(productionRate)
	return int(((rate * successRate) / 100) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	group := uint(((carsCount) / 10) * 95000)
	unit := uint((carsCount % 10) * 10000)

	return group + unit
}
