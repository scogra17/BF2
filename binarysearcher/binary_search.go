package binarysearcher

func BinarySearchFirstOccurence(target, indexLow, indexHigh int, inputArray []int) int {
	if indexHigh >= indexLow {
		indexMid := (indexLow + indexHigh)/2
		if inputArray[indexMid] == target && (indexMid == 0 || target > inputArray[indexMid - 1]) {
			return indexMid
		}
		if inputArray[indexMid] >= target {
			return BinarySearchFirstOccurence(target, indexLow, indexMid-1, inputArray)
		} else {
			return BinarySearchFirstOccurence(target, indexMid+1, indexHigh, inputArray)
		}
	}
	return -1
}

func BinarySearchLastOccurence(target, indexLow, indexHigh int, inputArray []int) int {
	if indexHigh >= indexLow {
		indexMid := (indexLow + indexHigh)/2
		if inputArray[indexMid] == target && (indexMid == 0 || target < inputArray[indexMid + 1]) {
			return indexMid
		}
		if inputArray[indexMid] <= target {
			return BinarySearchLastOccurence(target, indexMid+1, indexHigh, inputArray)
		} else {
			return BinarySearchLastOccurence(target, indexLow, indexMid-1, inputArray)
		}
	}
	return -1
}
