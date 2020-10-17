package tagetfinder

import "target_finder/binarysearch"

func FindValueInSliceBinary(target int, intSlice []int) int {
	targetFirstOccurrenceIndex := binarysearch.BinarySearchFirstOccurence(target, 0, len(intSlice)-1, intSlice)
	if targetFirstOccurrenceIndex == -1 {
		return -1
	}
	targetLastOccurrenceIndex := binarysearch.BinarySearchLastOccurence(target, 0, len(intSlice)-1, intSlice)

	numberOfOccurrences := targetLastOccurrenceIndex - targetFirstOccurrenceIndex + 1
	return numberOfOccurrences
}