package targetfinder

import (
	"target_finder/slicegenerator"
	"testing"
)

func benchmarkFindValueInSliceBinary(value int, slice []int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FindValueInSliceBinary(value, slice)
	}
}

func BenchmarkFindValueInSliceBinary10(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary100(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(100, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary1000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(1000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary10000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary100000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(100000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary1000000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(1000000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func BenchmarkFindValueInSliceBinary10000000(b *testing.B) {
	value, slice := slicegenerator.GenerateRandomOrderedSlice(10000000, 1000)
	benchmarkFindValueInSliceBinary(value, slice, b)
}

func TestFindValueInSliceBinary(t *testing.T) {
	testCases := []struct {
		array []int
		target int
		expectedCount int
	}{
		{
			array: []int{1,2,3,4},
			target: 3,
			expectedCount: 1,
		},
		{
			array: []int{1,2,3,3,3,3,3,4},
			target: 3,
			expectedCount: 5,
		},
		{
			array: []int{1,2,3,3,3,3,3,4},
			target: 5,
			expectedCount: 0,
		},
		{
			array: []int{},
			target: 1,
			expectedCount: 0,
		},

	}

	for _, test := range testCases {
		count := FindValueInSliceBinary(test.target, test.array)
		if count != test.expectedCount {
			t.Errorf("FindValueInSliceLinear failed, expected count of %d, got %d", test.expectedCount, count)
		}
	}
}