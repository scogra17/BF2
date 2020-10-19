package binarysearcher

// BinarySearchFirstOccurence finds the index of the first occurrence of target
// in inputArray
func BinarySearchFirstOccurence(target, indexLow, indexHigh int, intSlice []int) int {
	if indexHigh >= indexLow {
		indexMid := (indexLow + indexHigh)/2
		if intSlice[indexMid] == target && (indexMid == 0 || target > intSlice[indexMid - 1]) {
			return indexMid
		}
		if intSlice[indexMid] >= target {
			return BinarySearchFirstOccurence(target, indexLow, indexMid-1, intSlice)
		} else {
			return BinarySearchFirstOccurence(target, indexMid+1, indexHigh, intSlice)
		}
	}
	return -1
}

// BinarySearchFirstOccurence finds the index of the last occurrence of target
// in inputArray
func BinarySearchLastOccurence(target, indexLow, indexHigh int, intSlice []int) int {
	if indexHigh >= indexLow {
		indexMid := (indexLow + indexHigh)/2
		if intSlice[indexMid] == target && (indexMid == 0 || target < intSlice[indexMid + 1]) {
			return indexMid
		}
		if intSlice[indexMid] <= target {
			return BinarySearchLastOccurence(target, indexMid+1, indexHigh, intSlice)
		} else {
			return BinarySearchLastOccurence(target, indexLow, indexMid-1, intSlice)
		}
	}
	return -1
}
