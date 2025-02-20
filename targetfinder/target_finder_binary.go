package targetfinder

import "target_finder/binarysearcher"

// FindValueInSliceBinary uses a double binary search to find the number of
// occurrences of target in intSlice
func FindValueInSliceBinary(target int, intSlice []int) int {
	targetFirstOccurrenceIndex := binarysearcher.BinarySearchFirstOccurence(target, 0, len(intSlice)-1, intSlice)
	if targetFirstOccurrenceIndex == -1 {
		return 0
	}
	targetLastOccurrenceIndex := binarysearcher.BinarySearchLastOccurence(target, 0, len(intSlice)-1, intSlice)

	numberOfOccurrences := targetLastOccurrenceIndex - targetFirstOccurrenceIndex + 1
	return numberOfOccurrences
}