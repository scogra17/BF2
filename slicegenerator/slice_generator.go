package slicegenerator

import (
	"math/rand"
	"sort"
	"time"
)

// GenerateRandomOrderedSlice generates slice of ints of length length and with
// max value maxIntegerValue
func GenerateRandomOrderedSlice(length, maxIntegerValue int) (int, []int) {
	randomSlice := make([]int, length)
	seed := rand.NewSource(time.Now().UnixNano())
	randWithSeed := rand.New(seed)

	for i, _ := range randomSlice {
		randomSlice[i] = randWithSeed.Intn(maxIntegerValue)
	}
	sort.Ints(randomSlice)

	return randomSlice[randWithSeed.Intn(length-1)], randomSlice
}
