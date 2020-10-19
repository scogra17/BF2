package slicegenerator

import (
	"math/rand"
	"sort"
)

// GenerateRandomOrderedSlice generates slice of ints of length length and with
// max value maxIntegerValue
func GenerateRandomOrderedSlice(length, maxIntegerValue int) (int, []int) {
	randomSlice := make([]int, length)
	for i, _ := range randomSlice {
		randomSlice[i] = rand.Intn(maxIntegerValue)
	}
	sort.Ints(randomSlice)

	return randomSlice[rand.Intn(length-1)], randomSlice
}
