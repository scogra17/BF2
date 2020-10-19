package targetfinder

// FindValueInSliceLinear uses a linear scan to find the number of occurrences
// of target in intSlice
func FindValueInSliceLinear(target int, intSlice []int) int {
	targetCount := 0
	for _, value := range intSlice {
		if value == target {
			targetCount++
		}
	}
	return targetCount
}