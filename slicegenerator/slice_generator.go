package slicegenerator

import (
	"math/rand"
	"sort"
)


func GenerateRandomOrderedSlice(length, maxValue int) (int, []int) {
	randomSlice := make([]int, length)
	for i, _ := range randomSlice {
		randomSlice[i] = rand.Intn(maxValue)
	}
	sort.Ints(randomSlice)

	return randomSlice[rand.Intn(length-1)], randomSlice
}
