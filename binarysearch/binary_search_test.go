package binarysearch

import "testing"

func TestBinarySearchFirstOccurence(t *testing.T) {
	testCases := []struct {
		array []int
		target int
		expectedIndex int
 	}{
		{
			array: []int{1,2,3,4},
			target: 3,
			expectedIndex: 2,
		},
		{
			array: []int{1,2,3,3,3,3,3,4},
			target: 3,
			expectedIndex: 2,
		},
		{
			array: []int{},
			target: 1,
			expectedIndex: -1,
		},

	}

	for _, test := range testCases {
		index := BinarySearchFirstOccurence(test.target, 0,len(test.array)-1, test.array)
		if index != test.expectedIndex {
			t.Errorf("BinarySearch failed, expected index %d, go %d", test.expectedIndex, index)
		}
	}
}

func TestBinarySearchLastOccurence(t *testing.T) {
	testCases := []struct {
		array []int
		target int
		expectedIndex int
	}{
		{
			array: []int{1,2,3,4},
			target: 3,
			expectedIndex: 2,
		},
		{
			array: []int{1,2,3,3,3,3,3,4},
			target: 3,
			expectedIndex: 6,
		},
		{
			array: []int{},
			target: 1,
			expectedIndex: -1,
		},

	}

	for _, test := range testCases {
		index := BinarySearchLastOccurence(test.target, 0,len(test.array)-1, test.array)
		if index != test.expectedIndex {
			t.Errorf("BinarySearch failed, expected index %d, go %d", test.expectedIndex, index)
		}
	}
}

