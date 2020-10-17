package slicegenerator

import (
	"fmt"
	"testing"
)

func TestSliceGenerator(t *testing.T) {
	testCases := []struct{
		sliceLength int
		maxValue int
	}{
		{
			sliceLength: 9,
			maxValue: 100,
		},
	}

	for _, test := range testCases {
		_, slice := GenerateRandomOrderedSlice(test.sliceLength, test.maxValue)
		if len(slice) != test.sliceLength {
			t.Errorf("GenerateRandomOrderedSlice failed, expected slice of length %d, got %v", test.sliceLength, slice)
		}
		fmt.Printf("GenerateRandomOrderedSlice succeeded, expected slice of length %d, got %v", test.sliceLength, slice)
	}
}
