package tagetfinder

import "testing"

func TestFindValueInSliceLinear(t *testing.T) {
	slice1 := []int{1, 1, 2, 4, 5, 5, 7, 9}
	value1 := 5
	expectedCount := 2

	targetCount1 := FindValueInSliceLinear(value1, slice1)

	if targetCount1 != 2 {
		t.Errorf("TestFindValueInSliceLinear failed, expected %d, go %d", expectedCount, targetCount1)
	}
}
