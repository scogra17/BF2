package slicegenerator

import (
	"math/rand"
	"sort"
)


func GenerateRandomOrderedSlice(length, maxIntegerValue int) (int, []int) {
	randomSlice := make([]int, length)
	for i, _ := range randomSlice {
		randomSlice[i] = rand.Intn(maxIntegerValue)
	}
	sort.Ints(randomSlice)

	return randomSlice[rand.Intn(length-1)], randomSlice
}
