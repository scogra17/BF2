package slicegenerator

import (
	"testing"
)

func TestSliceGenerator(t *testing.T) {
	testCases := []struct{
		sliceLength int
		maxIntegerValue int
	}{
		{
			sliceLength: 9,
			maxIntegerValue: 100,
		},
	}

	for _, test := range testCases {
		_, slice := GenerateRandomOrderedSlice(test.sliceLength, test.maxIntegerValue)
		if len(slice) != test.sliceLength {
			t.Errorf("GenerateRandomOrderedSlice failed, expected slice of length %d, got %v", test.sliceLength, slice)
		}
	}
}
