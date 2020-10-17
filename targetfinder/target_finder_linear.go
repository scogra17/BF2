package targetfinder

func FindValueInSliceLinear(target int, intSlice []int) int {
	targetCount := 0
	for _, value := range intSlice {
		if value == target {
			targetCount++
		}
	}
	return targetCount
}