package main

import (
	"fmt"
	"target_finder/tagetfinder"
)

func main() {
	slice1 := []int{1, 1, 2, 4, 5, 5, 7, 9}
	value1 := 5
	targetCount1 := tagetfinder.FindValueInSliceLinear(value1, slice1)
	fmt.Println("targetCount1: %d", targetCount1)
}
